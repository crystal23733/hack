# Vulnerable Blog Application

이 프로젝트는 웹 보안 학습을 위한 의도적으로 취약점을 만든 블로그 애플리케이션입니다.
Go언어 공부도 할 겸 서버부분을 Go언어로 작성하였으나 Go언어의 템플릿이 기본적으로 이스케이프 처리가 되어 강제적으로 이스케이프처리하지 않도록 제작하였습니다.
학습에 필요할만한 부분 및 취약점은 지속적으로 업데이트할 예정입니다.

⚠️ **경고**: 이 애플리케이션은 교육 목적으로만 사용되어야 하며, 실제 운영 환경에 배포해서는 안 됩니다.

## 취약점 목록

이 애플리케이션은 다음과 같은 보안 취약점들을 포합하고 있습니다:

1. SQL Injection
    - 로그인 페이지
    - 게시글 상세 조회
    - 게시글 작성

2. Cross-Site Scripting (XSS)
    - 게시글 내용
    - 게시글 제목

## 시스템 요구사항

- Docker 및 Docker Compose (선택 사항)
- PostgreSQL (로컬 실행 시)
- Go 1.24 이상

## 설치 및 실행 방법

### Docker Compose 사용

1. 저장소 클론
```bash
git clone https://github.com/crystal23733/hack.git
cd vulnerable-blog
```

2. 애플리케이션 실행
```bash
sudo docker compose up -d --build
```

3. 브라우저 접속
http://localhost

### 로컬 환경에서 실행

1. PostgreSQL 설치 및 실행

2. 데이터베이스 초기화
```bash
psql -U postgres -f database/init.sql
```

3. Go 애플리케이션 실행
```bash
cd client
go mod download
go run main.go
```

4. 브라우저에서 접속
http://localhost:8080

## 기본 계정 정보

- 사용자명: admin
- 비밀번호: password123

## 취약점 테스트 예시

### SQL Injection

1. 로그인 우회:
```sql
'or''='
```

2. 게시글 조회 시 인젝션(예 - 다중 게시글 인젝션):
```sql
test', 'test'); INSERT INTO posts (user_id, title, content) VALUES (1, '해킹1', '내용1'), (1, '해킹2', '내용2'); --
```

### XSS (Cross-Site Scriping)

게시글 작성 시 다음과 같은 내용 비력:
```html
<script>alert('XSS 공격 실행')</script>
```

## 주의사항

1. 이 애플리케이션은 학습 목적으로만 사용해야 합니다.
2. 로컬 환경이나 격리된 환경에서만 실행하세요.
3. 실제 개인정보나 중요 데이터를 사용하지 마세요.
4. 다른 사람의 시스템에 대해 이 취약점들을 테스트하지 마세요

## 디렉토리 구조

```bash
vulnerable-blog/
├── client/ # Go 웹 애플리케이션
│ ├── static/ # 정적 파일 (HTML, CSS, JS)
│ ├── service/ # 핵심 비즈니스 로직
│ └── Dockerfile
├── database/ # PostgreSQL 설정
│ ├── init.sql # 초기 데이터베이스 스키마
│ └── Dockerfile
├── nginx/ # Nginx 설정
│ ├── nginx.conf
│ └── Dockerfile
└── docker-compose.yml # Docker Compose 설정
```

## 기여하기

버그 리포트나 기능 개선 제안은 GitHub 이슈를 통해 제출해주세요.