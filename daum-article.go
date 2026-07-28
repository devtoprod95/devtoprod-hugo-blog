package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// NewsArticle represents a scraped news article
type NewsArticle struct {
	ID          string
	Title       string
	Category    string
	Author      string
	Date        time.Time
	Content     string
	OriginalURL string
}

func main() {
	fmt.Println("=== Daum Entertainment News Scraper Start ===")

	// 1. Get recent article URLs from the entertainment news list page
	mainURL := "https://entertain.daum.net/news"
	html, err := fetchHTML(mainURL)
	if err != nil {
		fmt.Printf("Failed to fetch main page: %v\n", err)
		return
	}

	urls := extractArticleURLs(html)
	if len(urls) == 0 {
		fmt.Println("No article URLs found from the main page.")
		return
	}
	fmt.Printf("Found %d unique article URLs. Starting categorization and scraping...\n", len(urls))

	// Target categories we want to collect (1 article per category)
	targetCategories := map[string]string{
		"드라마":  "drama",
		"예능":   "variety",
		"영화":   "movie",
		"뮤직":   "music",
		"해외연예": "foreign",
	}

	collected := make(map[string]bool)

	for _, url := range urls {
		// Stop if we have collected all 5 categories
		if len(collected) == len(targetCategories) {
			break
		}

		articleID := extractArticleID(url)
		if articleID == "" {
			continue
		}

		// Scrape detailed page
		detailHTML, err := fetchHTML(url)
		if err != nil {
			fmt.Printf("[%s] Failed to fetch article details: %v\n", articleID, err)
			continue
		}

		article, err := parseArticle(articleID, detailHTML, url)
		if err != nil {
			// Skip if it fails or if it's not in our target categories
			continue
		}

		// Map Daum category name to our target categories
		var mappedKey string
		switch article.Category {
		case "드라마":
			mappedKey = "드라마"
		case "예능":
			mappedKey = "예능"
		case "영화":
			mappedKey = "영화"
		case "가요", "뮤직":
			mappedKey = "뮤직"
		case "해외연예":
			mappedKey = "해외연예"
		default:
			// Unmapped categories will be ignored to focus on user's requested scope
			continue
		}

		// Skip if we already collected an article for this category
		if collected[mappedKey] {
			continue
		}

		// Save article to markdown file
		err = saveToMarkdown(article, mappedKey, targetCategories[mappedKey])
		if err != nil {
			fmt.Printf("[%s] Failed to save article: %v\n", articleID, err)
			continue
		}

		fmt.Printf("Successfully generated article! Category: [%s], Title: %s\n", mappedKey, article.Title)
		collected[mappedKey] = true
	}

	fmt.Println("=== Daum Entertainment News Scraper Finish ===")
}

// fetchHTML fetches HTML from a URL with a fake User-Agent
func fetchHTML(targetURL string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return "", err
	}

	// Set header to prevent bot blocking
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// extractArticleURLs parses unique Daum article URLs from HTML body
func extractArticleURLs(html string) []string {
	re := regexp.MustCompile(`https://v\.daum\.net/v/\d+`)
	matches := re.FindAllString(html, -1)

	// Remove duplicates
	uniqueMap := make(map[string]bool)
	var urls []string
	for _, m := range matches {
		if !uniqueMap[m] {
			uniqueMap[m] = true
			urls = append(urls, m)
		}
	}
	return urls
}

// extractArticleID extracts numerical ID from URL
func extractArticleID(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

// parseArticle parses article information from detail HTML
func parseArticle(id, html, originalURL string) (NewsArticle, error) {
	// 1. Parse Category (<h2 class="screen_out">영화</h2>)
	categoryRe := regexp.MustCompile(`<h2 class="screen_out">([^<]+)</h2>`)
	categoryMatch := categoryRe.FindStringSubmatch(html)
	category := "일반"
	if len(categoryMatch) > 1 {
		category = strings.TrimSpace(categoryMatch[1])
	}

	// 2. Parse Title
	titleRe := regexp.MustCompile(`<h3 class="tit_view"[^>]*>([^<]+)</h3>`)
	titleMatch := titleRe.FindStringSubmatch(html)
	if len(titleMatch) < 2 {
		return NewsArticle{}, fmt.Errorf("title not found")
	}
	title := strings.TrimSpace(titleMatch[1])
	// Escape HTML special characters
	title = htmlUnescape(title)

	// 3. Parse Author
	authorRe := regexp.MustCompile(`<span class="txt_info">([^<]+)</span>`)
	authorMatches := authorRe.FindAllStringSubmatch(html, -1)
	author := "연예 뉴스 기자"
	if len(authorMatches) > 0 && len(authorMatches[0]) > 1 {
		author = strings.TrimSpace(authorMatches[0][1])
	}

	// 4. Parse Date (format: "2026. 7. 28. 16:31")
	dateRe := regexp.MustCompile(`<span class="num_date">([0-9.]+)\s*([0-9:]+)</span>`)
	dateMatch := dateRe.FindStringSubmatch(html)
	parsedDate := time.Now()
	if len(dateMatch) >= 3 {
		dateStr := fmt.Sprintf("%s %s", strings.TrimSpace(dateMatch[1]), strings.TrimSpace(dateMatch[2]))
		// clean up trailing dots: "2026. 7. 28. 16:31" -> "2026. 7. 28. 16:31"
		dateStr = strings.ReplaceAll(dateStr, " ", "")
		
		// Regex parse details
		parseParts := regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)\.(\d+):(\d+)`).FindStringSubmatch(dateStr)
		if len(parseParts) >= 6 {
			year := parseParts[1]
			month := fmt.Sprintf("%02s", parseParts[2])
			day := fmt.Sprintf("%02s", parseParts[3])
			hour := fmt.Sprintf("%02s", parseParts[4])
			minute := fmt.Sprintf("%02s", parseParts[5])
			
			stdFormat := fmt.Sprintf("%s-%s-%sT%s:%s:00+09:00", year, month, day, hour, minute)
			t, err := time.Parse("2006-01-02T15:04:05-07:00", stdFormat)
			if err == nil {
				parsedDate = t
			}
		}
	}

	// 5. Parse Content (<p dmcf-ptype="general">...)
	pRe := regexp.MustCompile(`<p[^>]*dmcf-ptype="general"[^>]*>(.*?)</p>`)
	pMatches := pRe.FindAllStringSubmatch(html, -1)
	
	var contentBuilder strings.Builder
	tagRe := regexp.MustCompile(`<[^>]*>`) // For stripping nested HTML tags inside paragraphs
	
	for _, match := range pMatches {
		if len(match) > 1 {
			pText := match[1]
			// Remove HTML Tags
			pText = tagRe.ReplaceAllString(pText, "")
			pText = htmlUnescape(strings.TrimSpace(pText))
			if pText != "" {
				contentBuilder.WriteString(pText)
				contentBuilder.WriteString("\n\n")
			}
		}
	}

	content := contentBuilder.String()
	if content == "" {
		return NewsArticle{}, fmt.Errorf("content is empty")
	}

	return NewsArticle{
		ID:          id,
		Title:       title,
		Category:    category,
		Author:      author,
		Date:        parsedDate,
		Content:     content,
		OriginalURL: originalURL,
	}, nil
}

// htmlUnescape replaces common HTML entities
func htmlUnescape(raw string) string {
	r := strings.NewReplacer(
		"&quot;", "\"",
		"&amp;", "&",
		"&lt;", "<",
		"&gt;", ">",
		"&#39;", "'",
		"&nbsp;", " ",
		"&middot;", "·",
	)
	return r.Replace(raw)
}

// saveToMarkdown saves the article as a Hugo Markdown post
func saveToMarkdown(article NewsArticle, koreanCat, englishCat string) error {
	// Target Directory: blog/content/posts/news/entertain/<englishCat>/
	dirPath := filepath.Join("blog", "content", "posts", "news", "entertain", englishCat)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s.md", article.ID)
	filePath := filepath.Join(dirPath, fileName)

	// Check if already exists (skip duplicates)
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("[%s] Article already exists on local blog. Skipping.\n", article.ID)
		return nil
	}

	// Build markdown template
	mdContent := fmt.Sprintf(`---
title: "%s"
date: %s
draft: false
author: "%s"
originalUrl: "%s"
categories: ["기사", "엔터", "%s"]
---

%s

---
> **출처**: 본 기사는 Daum 연예 기사에서 제공된 뉴스 정보입니다. 원본 기사는 [이곳](%s)에서 확인하실 수 있습니다. 저작권은 해당 언론사 및 기자에게 있습니다.
`, 
		strings.ReplaceAll(article.Title, "\"", "\\\""),
		article.Date.Format("2006-01-02T15:04:05+09:00"),
		article.Author,
		article.OriginalURL,
		koreanCat,
		article.Content,
		article.OriginalURL,
	)

	return os.WriteFile(filePath, []byte(mdContent), 0644)
}
