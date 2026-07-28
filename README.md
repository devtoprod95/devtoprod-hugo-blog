# 🚀 Go + Hugo 블로그 프로젝트 (go-hugo-blog)

이 프로젝트는 **Go 자동화 스크립트**와 **Hugo 정적 사이트 생성기(Static Site Generator)**를 **Docker Compose** 개발 환경으로 통합한 블로그 프로젝트입니다. 로컬에서 실시간으로 글을 작성 및 확인하고, GitHub Actions를 통해 깃허브 페이지(GitHub Pages)에 원클릭으로 배포할 수 있도록 구성되어 있습니다.

---

## 📌 주요 특징 및 기능

1. **실시간 프리뷰 (Live Reload)**: 로컬 컨테이너 환경에서 글을 작성하거나 설정을 바꾸면 1초 내로 브라우저에 바로 반영됩니다.
2. **포스팅 생성 자동화**: Go 언어로 작성된 스크립트가 실행될 때마다 Front Matter 양식을 만족하는 날짜 기반의 테스트 포스팅을 자동 생성해 줍니다.
3. **강력한 로컬 개발 분기**: 개발 중인 로컬 도메인(`http://localhost:1313`)과 실제 배포 도메인이 설정 파일 충돌 없이 안전하게 분리 작동합니다.
4. **CI/CD 자동 배포**: `master` 브랜치에 코드를 푸시하기만 하면 GitHub Actions가 자동으로 빌드하여 `gh-pages` 브랜치로 배포합니다.

---

## 🛠️ 기술 스택 (Tech Stack)

* **Static Site Generator**: Hugo v0.154.5 (Extended version)
* **Blog Theme**: `seotax` (Advanced Taxonomy Search & Multi-language support)
* **Script Engine**: Go 1.22-alpine
* **Orchestration**: Docker Compose
* **CI/CD**: GitHub Actions

---

## 📂 프로젝트 구조 (Directory Structure)

```text
├── .github/workflows/
│   └── deploy.yml          # GitHub Actions 자동 배포 워크플로우
├── blog/
│   ├── content/
│   │   └── posts/          # 블로그 포스트 마크다운 파일 저장 위치
│   ├── static/
│   │   └── images/         # 이미지, 파비콘, 프로필 사진 등의 정적 자원
│   ├── themes/
│   │   └── seotax/         # 블로그 테마 레이아웃 및 스타일 소스
│   └── config.yaml         # 블로그 기본 메타데이터 및 메뉴 설정 파일
├── Dockerfile.go           # Go 포스팅 자동 생성기 컨테이너용 도커파일
├── docker-compose.yml      # Hugo 및 Go 앱 로컬 개발 환경 오케스트레이션
├── go.mod                  # Go 모듈 의존성 설정
└── main.go                 # 자동 포스팅 생성기 Go 소스 코드
```

---

## 💻 로컬 개발 환경 실행 방법

### 1. 전체 개발 환경 실행 (Docker Compose)
터미널에서 아래 명령을 실행하여 로컬 Hugo 서버와 Go 자동 포스팅 생성기를 실행합니다.
```bash
docker compose up -d
```
* **블로그 접속**: 웹 브라우저에서 [http://localhost:1313](http://localhost:1313)으로 이동하여 실시간 반영 결과를 확인합니다.
* **로컬 캐시 초기화 방지**: 로컬 개발 서버 구동 명령어에는 `--ignoreCache` 및 `--disableFastRender` 옵션이 들어가 있어, 도커 재기동 없이도 카테고리나 글 수정이 즉시 동기화됩니다.

### 2. 포스팅 생성기 단독 실행 (로컬 Go 환경)
도커를 띄우지 않고 로컬에 설치된 Go 환경에서 새 테스트 글을 생성하고 싶다면 아래 명령을 사용합니다.
```bash
go run main.go
```
실행 결과로 `blog/content/posts/YYYY-MM-DD-HHMMSS-docker-test.md` 형태의 마크다운 파일이 자동 생성됩니다.

---

## 🚀 깃허브 페이지 배포 방법 (GitHub Pages Deployment)

1. **깃허브 저장소 업로드**:
   작업이 완료된 소스 코드를 본인의 깃허브 저장소(예: `devtoprod95/devtoprod-hugo-blog`)의 `master` 브랜치로 푸시합니다.
   ```bash
   git add .
   git commit -m "feat: update blog content"
   git push origin master
   ```
2. **자동 빌드 & 배포**:
   푸시와 동시에 [GitHub Actions 워크플로우](.github/workflows/deploy.yml)가 트리거되어 자동으로 빌드를 완료하고 `gh-pages` 브랜치를 생성하여 결과물을 업로드합니다.
3. **GitHub Pages 서비스 활성화**:
   * 본인의 깃허브 저장소 웹페이지로 접속하여 **`Settings` ➡️ `Pages`** 메뉴로 이동합니다.
   * **Build and deployment** 섹션의 Branch 항목에서 **`gh-pages`** 브랜치와 `/ (root)` 폴더를 선택하고 **`Save`**를 누르면 배포 완료됩니다.
   * 배포 완료 후 블로그 주소: `https://devtoprod95.github.io/devtoprod-hugo-blog/`

---

## ⚖️ 라이선스 및 출처 고지 (Credits)

이 블로그의 디자인과 레이아웃은 [minyeamer의 hugo-seotax 테마](https://github.com/minyeamer/hugo-seotax)를 기반으로 커스텀하여 제작되었습니다. 

해당 테마는 **MIT 라이선스** 하에 제공되며, 원저작자 표기 및 라이선스 고지 의무를 준수합니다.

* **원저작자 (Original Author)**: [minyeamer](https://github.com/minyeamer)
* **원본 저장소 (Original Repository)**: [hugo-seotax](https://github.com/minyeamer/hugo-seotax)
* **라이선스 (License)**: MIT License (Copyright (c) minyeamer)
