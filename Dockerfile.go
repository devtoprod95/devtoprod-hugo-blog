FROM golang:1.22-alpine

WORKDIR /app

# 기본 패키지 설치
RUN apk add --no-cache git tzdata

COPY go.mod ./
RUN go mod download || true

COPY . .

CMD ["go", "run", "main.go"]