package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// ============================================================================
// ⚠️ 중요 안내
// 이 코드를 작성하는 시점에 m.entertain.naver.com 페이지에 직접 접속해서
// 실제 HTML 구조를 확인할 수 없었습니다 (네트워크 접근 제한).
// 따라서 아래 categories의 URL, 그리고 기사 링크를 찾는 정규식은 "추정"입니다.
// 실행 전에 반드시:
//   1. 브라우저에서 https://m.entertain.naver.com/home 접속
//   2. 상단 카테고리 탭(드라마/영화/방송/화보/인기/연재 등) 클릭해서 실제 URL 확인
//   3. categories 슬라이스의 URL을 실제 값으로 교체
//   4. 개발자 도구(F12) > Elements에서 기사 링크 <a href="..."> 패턴 확인 후
//      articleLinkPattern 정규식이 맞는지 확인
// 기사 제목/요약은 CSS 선택자 대신 OG 메타태그(og:title, og:description)를
// 우선 사용합니다. 이는 페이지 디자인이 바뀌어도 잘 안 깨지는 안정적인 방법입니다.
// ============================================================================

// Category 카테고리 정보
type Category struct {
	Name string // 화면/파일명에 쓸 한글 카테고리명
	Slug string // 파일명/태그용 영문 slug
	URL  string // 카테고리별 목록 페이지 URL (★ 검증 필요)
}

// 실제 사이트에서 확인된 카테고리 URL입니다.
var categories = []Category{
	{Name: "드라마", Slug: "drama", URL: "https://m.entertain.naver.com/drama"},
	{Name: "영화", Slug: "movie", URL: "https://m.entertain.naver.com/movie"},
	{Name: "뮤직", Slug: "music", URL: "https://m.entertain.naver.com/music"},
	{Name: "연예", Slug: "relationship", URL: "https://m.entertain.naver.com/relationship"},
	{Name: "포토", Slug: "pictorial", URL: "https://m.entertain.naver.com/photo"},
	{Name: "랭킹", Slug: "ranking", URL: "https://m.entertain.naver.com/ranking"},
}

// ★ 검증 필요: 네이버 엔터 기사 URL은 전통적으로 /article/{oid}/{aid} 형태를 사용합니다.
// 다른 패턴(예: /now/... , ?oid=...&aid=...)일 수도 있으니 실제 링크를 보고 조정하세요.
var articleLinkPattern = regexp.MustCompile(`entertain\.naver\.com/.*article/(\d+)/(\d+)`)

const (
	postsDirRel   = "blog/content/posts"
	stateDirRel   = "blog/.state"
	stateFileName = "naver-collected.json"
	userAgent     = "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1"
	requestDelay  = 800 * time.Millisecond
)

// CollectedState 이미 수집한 기사 URL을 기록해서 중복 수집을 방지
type CollectedState struct {
	CollectedURLs map[string]string `json:"collected_urls"` // url -> 수집일시
}

func loadState(path string) (*CollectedState, error) {
	state := &CollectedState{CollectedURLs: map[string]string{}}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return state, nil
		}
		return nil, err
	}
	if err := json.Unmarshal(data, state); err != nil {
		return nil, err
	}
	if state.CollectedURLs == nil {
		state.CollectedURLs = map[string]string{}
	}
	return state, nil
}

func saveState(path string, state *CollectedState) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// browser는 프로그램 전체에서 재사용하는 헤드리스 브라우저 인스턴스입니다.
// (요청마다 브라우저를 새로 띄우면 느리고 리소스 낭비가 심하므로 공유합니다)
var browser *rod.Browser

// initBrowser 헤드리스 브라우저를 띄워서 browser 전역 변수를 초기화한다.
// 반환된 cleanup 함수는 main()에서 defer로 반드시 호출해서 브라우저를 정리해야 한다.
func initBrowser() (func(), error) {
	execPath := os.Getenv("CHROME_BIN") // Dockerfile에서 지정한 /usr/bin/chromium-browser 등

	l := launcher.New().
		Headless(true).
		Set("no-sandbox").              // 컨테이너 환경(주로 root 실행)에서 필요
		Set("disable-dev-shm-usage").   // /dev/shm 용량이 작은 컨테이너 대응
		Set("disable-crash-reporter").  // non-root 컨테이너에서 crashpad 에러 방지
		Set("disable-gpu")
	if execPath != "" {
		l = l.Bin(execPath)
	}

	controlURL, err := l.Launch()
	if err != nil {
		return nil, fmt.Errorf("헤드리스 브라우저 실행 실패 (CHROME_BIN=%q 확인 필요): %w", execPath, err)
	}

	b := rod.New().ControlURL(controlURL)
	if err := b.Connect(); err != nil {
		return nil, fmt.Errorf("브라우저 연결 실패: %w", err)
	}

	browser = b
	return func() {
		b.Close()
		l.Cleanup()
	}, nil
}

// fetchDoc 헤드리스 브라우저로 페이지를 열어 JS 렌더링이 끝난 뒤의 HTML을 가져온다.
// go-rod의 WaitStable은 DOM 변화와 네트워크 요청이 잠잠해질 때까지 자동으로
// 기다려주므로, chromedp에서처럼 고정 Sleep 시간을 추측해서 넣지 않아도 된다.
func fetchDoc(url string) (*goquery.Document, error) {
	page, err := browser.Timeout(45 * time.Second).Page(proto.TargetCreateTarget{})
	if err != nil {
		return nil, fmt.Errorf("페이지 생성 실패: %w", err)
	}
	defer page.Close()

	// m.entertain.naver.com은 모바일 전용 페이지이므로 모바일 UA로 접속해야
	// 데스크톱용으로 리다이렉트되지 않는다.
	if err := page.SetUserAgent(&proto.NetworkSetUserAgentOverride{UserAgent: userAgent}); err != nil {
		return nil, fmt.Errorf("User-Agent 설정 실패: %w", err)
	}

	if err := page.Navigate(url); err != nil {
		return nil, fmt.Errorf("페이지 이동 실패: %w", err)
	}

	if err := page.WaitLoad(); err != nil {
		return nil, fmt.Errorf("페이지 로드 대기 실패: %w", err)
	}

	// DOM/네트워크가 안정될 때까지 대기. SPA가 계속 폴링 요청을 보내는 경우
	// 타임아웃될 수 있는데, 그 경우엔 에러를 무시하고 현재 HTML을 그대로 사용한다.
	if err := page.WaitStable(1 * time.Second); err != nil {
		fmt.Printf("  (참고) WaitStable 타임아웃, 현재 상태로 진행합니다: %v\n", err)
	}

	html, err := page.HTML()
	if err != nil {
		return nil, fmt.Errorf("HTML 가져오기 실패: %w", err)
	}

	return goquery.NewDocumentFromReader(strings.NewReader(html))
}

// fetchDocWithRetry는 fetchDoc이 실패하면(특히 첫 요청의 콜드 스타트 타임아웃)
// 짧게 대기 후 한 번 더 시도한다.
func fetchDocWithRetry(url string) (*goquery.Document, error) {
	doc, err := fetchDoc(url)
	if err == nil {
		return doc, nil
	}
	fmt.Printf("  ⏳ 1차 요청 실패(%v), 3초 후 재시도...\n", err)
	time.Sleep(3 * time.Second)
	return fetchDoc(url)
}

// findFirstArticleURL 카테고리 목록 페이지에서 아직 수집하지 않은
// 첫 번째 기사 URL을 찾는다.
func findFirstArticleURL(doc *goquery.Document, state *CollectedState) (string, bool) {
	var found string
	doc.Find("a[href]").EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, _ := s.Attr("href")
		if href == "" {
			return true // continue
		}
		if !strings.HasPrefix(href, "http") {
			// 상대경로인 경우 도메인 붙이기
			href = "https://m.entertain.naver.com" + href
		}
		if articleLinkPattern.MatchString(href) {
			if _, already := state.CollectedURLs[href]; already {
				return true // 이미 수집한 기사, 다음 링크 확인
			}
			found = href
			return false // 중단
		}
		return true
	})
	return found, found != ""
}

// debugPrintLinks 목록 페이지에서 기사를 못 찾았을 때, 실제로 페이지에
// 어떤 링크들이 있는지 보여줘서 URL 패턴/렌더링 방식을 진단할 수 있게 한다.
func debugPrintLinks(doc *goquery.Document) {
	total := 0
	shown := 0
	var candidateLinks []string // "숫자 id를 포함한" 콘텐츠성 링크로 추정되는 것들

	// 숫자가 포함된 경로 세그먼트가 있으면 개별 콘텐츠(기사/사진/영상 등)일 가능성이 높다.
	numericSegment := regexp.MustCompile(`/\d{5,}`)

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		total++
		href, _ := s.Attr("href")

		if shown < 15 {
			text := strings.TrimSpace(s.Text())
			if len(text) > 30 {
				text = text[:30] + "..."
			}
			fmt.Printf("      href=%s  text=%q\n", href, text)
			shown++
		}

		if strings.Contains(href, "naver.com") && numericSegment.MatchString(href) {
			candidateLinks = append(candidateLinks, href)
		}
	})
	fmt.Printf("      (총 <a href> %d개 중 %d개 표시)\n", total, shown)

	if len(candidateLinks) > 0 {
		fmt.Printf("\n      💡 숫자 ID가 포함된 콘텐츠성 링크 후보 (최대 10개):\n")
		limit := 10
		if len(candidateLinks) < limit {
			limit = len(candidateLinks)
		}
		for _, l := range candidateLinks[:limit] {
			fmt.Printf("         %s\n", l)
		}
		fmt.Printf("      ⬆️ 이 중 실제 기사/사진 링크로 보이는 URL 패턴을 알려주시면 정규식을 맞춰드릴 수 있습니다.\n")
	}

	if total == 0 {
		fmt.Printf("      ⚠️  페이지에 <a> 태그가 전혀 없습니다. JS로 렌더링되는 SPA 페이지일 가능성이 높습니다.\n")
		fmt.Printf("         (이 경우 goquery만으로는 못 가져오고, 헤드리스 브라우저나\n")
		fmt.Printf("          해당 페이지가 호출하는 내부 API(JSON)를 개발자 도구 Network 탭에서 찾아야 합니다.)\n")
	}

	// 참고용으로 og:title도 출력 (페이지 자체는 정상 로드됐는지 확인)
	pageTitle := doc.Find(`meta[property="og:title"]`).AttrOr("content", "")
	if pageTitle == "" {
		pageTitle = doc.Find("title").First().Text()
	}
	fmt.Printf("      페이지 제목(og:title/title): %q\n", strings.TrimSpace(pageTitle))
}

// ArticleInfo 기사에서 추출한 정보
type ArticleInfo struct {
	Title   string
	Summary string
	Image   string
	URL     string
}

// extractArticleInfo OG 메타태그를 우선으로 기사 정보를 추출한다.
// (CSS 클래스는 자주 바뀌지만 OG 태그는 SEO 목적상 잘 안 바뀌는 편)
func extractArticleInfo(doc *goquery.Document, url string) ArticleInfo {
	info := ArticleInfo{URL: url}

	info.Title = doc.Find(`meta[property="og:title"]`).AttrOr("content", "")
	info.Summary = doc.Find(`meta[property="og:description"]`).AttrOr("content", "")
	info.Image = doc.Find(`meta[property="og:image"]`).AttrOr("content", "")

	// og:title이 비어있으면 <title> 태그로 폴백
	if info.Title == "" {
		info.Title = strings.TrimSpace(doc.Find("title").First().Text())
	}

	// ★ 선택: 본문 일부를 더 가져오고 싶다면 아래 선택자를 실제 구조 확인 후 사용
	// (본문 전체를 그대로 재게시하는 것은 저작권상 바람직하지 않으므로
	//  요약/발췌 목적으로만 짧게 사용하는 것을 권장합니다)
	body := doc.Find("#dic_area, #articeBody, .article_body").First().Text()
	info.Summary = strings.TrimSpace(body)
	
	return info
}

func slugify(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 60 {
		s = s[:60]
	}
	replacer := strings.NewReplacer(" ", "-", "\"", "", "'", "", "/", "-", "\n", "")
	return replacer.Replace(s)
}

func writePost(category Category, info ArticleInfo) (string, error) {
	if err := os.MkdirAll(postsDirRel, 0755); err != nil {
		return "", fmt.Errorf("디렉터리 생성 실패: %w", err)
	}

	now := time.Now()
	fileName := fmt.Sprintf("%s-naver-%s.md", now.Format("2006-01-02-150405"), category.Slug)
	filePath := filepath.Join(postsDirRel, fileName)

	title := info.Title
	if title == "" {
		title = fmt.Sprintf("[%s] 네이버 엔터 기사", category.Name)
	}

	summary := info.Summary
	if summary == "" {
		summary = "요약 정보를 가져오지 못했습니다. 원문 링크를 확인해주세요."
	}

	content := fmt.Sprintf(`---
title: %q
date: %s
description: %q
draft: false
tags: ["엔터", %q]
categories: ["엔터", %q]
source_url: %q
---

## %s

%s

---

- **카테고리**: %s
- **원문 링크**: [%s](%s)
- **수집 시각**: %s
`,
		title,
		now.Format(time.RFC3339),
		summary,
		category.Name,
		category.Name,
		info.URL,
		title,
		summary,
		category.Name,
		info.URL,
		info.URL,
		now.Format("2006-01-02 15:04:05"),
	)

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("파일 생성 실패: %w", err)
	}

	return filePath, nil
}

func main() {
	fmt.Println("🚀 네이버 엔터 기사 수집 시작...")

	fmt.Println("🌐 헤드리스 브라우저 초기화 중...")
	cleanup, err := initBrowser()
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		return
	}
	defer cleanup()

	if err := os.MkdirAll(stateDirRel, 0755); err != nil {
		fmt.Printf("❌ 상태 디렉터리 생성 실패: %v\n", err)
		return
	}
	statePath := filepath.Join(stateDirRel, stateFileName)

	state, err := loadState(statePath)
	if err != nil {
		fmt.Printf("❌ 상태 파일 로드 실패: %v\n", err)
		return
	}

	successCount := 0

	for _, category := range categories {
		fmt.Printf("\n📂 [%s] 카테고리 처리 중... (%s)\n", category.Name, category.URL)

		listDoc, err := fetchDocWithRetry(category.URL)
		if err != nil {
			fmt.Printf("  ⚠️  목록 페이지 요청 실패: %v (URL 확인 필요)\n", err)
			continue
		}

		articleURL, ok := findFirstArticleURL(listDoc, state)
		if !ok {
			fmt.Printf("  ⚠️  새 기사를 찾지 못했습니다. 디버그를 위해 발견된 링크들을 출력합니다:\n")
			debugPrintLinks(listDoc)
			continue
		}

		time.Sleep(requestDelay) // 서버 부담을 줄이기 위한 딜레이

		articleDoc, err := fetchDocWithRetry(articleURL)
		if err != nil {
			fmt.Printf("  ⚠️  기사 페이지 요청 실패: %v\n", err)
			continue
		}

		info := extractArticleInfo(articleDoc, articleURL)

		filePath, err := writePost(category, info)
		if err != nil {
			fmt.Printf("  ❌ 파일 생성 실패: %v\n", err)
			continue
		}

		state.CollectedURLs[articleURL] = time.Now().Format(time.RFC3339)
		if err := saveState(statePath, state); err != nil {
			fmt.Printf("  ⚠️  상태 저장 실패: %v\n", err)
		}

		fmt.Printf("  ✅ 생성 완료: %s\n", filePath)
		fmt.Printf("     제목: %s\n", info.Title)
		fmt.Printf("     원문: %s\n", articleURL)

		successCount++
		time.Sleep(requestDelay)
	}

	fmt.Printf("\n🎉 완료: %d/%d 카테고리에서 기사 수집 성공\n", successCount, len(categories))
}