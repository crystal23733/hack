# 빌드 스테이지
FROM golang:1.24-alpine AS builder

# 작업 디렉토리 설정
WORKDIR /app

# 필요한 시스템 패키지 설치
RUN apk add --no-cache gcc musl-dev

# Go 모듈 파일 복사 및 의존성 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사
COPY . .

# 애플리케이션 빌드
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# 실행 스테이지
FROM alpine:latest

# 작업 디렉토리 설정
WORKDIR /app

# 빌드 스테이지에서 빌드된 바이너리 복사
COPY --from=builder /app/main .
# static 폴더 복사 (HTML, CSS, JavaScript 파일들)
COPY --from=builder /app/static ./static

# 포트 설정
EXPOSE 8080

# 애플리케이션 실행
CMD ["./main"]