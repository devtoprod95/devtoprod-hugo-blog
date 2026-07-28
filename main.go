package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("🚀 Go 스크립트 실행: 테스트 게시글 생성 시작...")

	// 1. 파일 저장 경로 지정 (blog/content/posts/)
	postsDir := filepath.Join("blog", "content", "posts")
	
	// posts 디렉터리가 없으면 생성
	if err := os.MkdirAll(postsDir, 0755); err != nil {
		fmt.Printf("❌ 디렉터리 생성 실패: %v\n", err)
		return
	}

	// 2. 파일명 및 Front Matter(Hugo 메타데이터) 설정
	now := time.Now()
	fileName := fmt.Sprintf("%s-docker-test.md", now.Format("2006-01-02-150405"))
	filePath := filepath.Join(postsDir, fileName)

	fileContent := fmt.Sprintf(`---
title: "Docker + Go로 생성된 자동 포스팅 테스트"
date: %s
draft: false
tags: ["Go", "Hugo", "Docker"]
---

## 📌 안녕! Go + Hugo 연동 성공!

이 글은 **Go 스크립트**가 도커 환경에서 직접 생성한 마크다운 파일입니다.

### 테스트 정보
- **생성 시각**: %s
- **배포 방식**: Docker Compose + Hugo Live Reload
- **상태**: 정상 동작 확인 완료
`, now.Format(time.RFC3339), now.Format("2006-01-02 15:04:05"))

	// 3. 마크다운 파일 쓰기
	err := os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		fmt.Printf("❌ 파일 생성 실패: %v\n", err)
		return
	}

	fmt.Printf("✅ 마크다운 생성 완료: %s\n", filePath)
}