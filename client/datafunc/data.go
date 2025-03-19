package datafunc

import (
	"database/sql"
	"fmt"
	"hack/client/config"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Data() {
	// 환경변수 로드
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("환경변수 로드 실패", err)
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_USER,
		cfg.DB_PASSWORD,
		cfg.DB_NAME,
	)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("데이터베이스 연결 실패", err)
	}

	// 데이터베이스 연결 확인
	err = DB.Ping()
	if err != nil {
		log.Fatal("데이터베이스 연결 확인 실패", err)
	}

	log.Println("데이터베이스 연결 성공")

}
