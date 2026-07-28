FROM golang:1.22-alpine

WORKDIR /app

# 기본 패키지 + go-rod가 쓸 Chromium과 headless 렌더링에 필요한 폰트/라이브러리
RUN apk add --no-cache \
    git \
    tzdata \
    chromium \
    nss \
    freetype \
    freetype-dev \
    harfbuzz \
    ca-certificates \
    ttf-freefont

# go-rod가 크롬 실행 파일을 찾을 수 있도록 지정
# (alpine chromium 패키지는 보통 /usr/bin/chromium-browser 에 설치됨)
ENV CHROME_BIN=/usr/bin/chromium-browser

COPY go.mod go.sum ./
RUN go mod download || true

COPY . .

# CMD는 프로젝트 상황(다른 진입점과의 충돌 등)에 맞게 직접 지정해주세요.
# 예: CMD ["go", "run", "naver-article.go"]
# 또는 컨테이너를 띄워둔 뒤 필요할 때만 수동 실행:
#   docker compose run --rm <서비스명> go run naver-article.go