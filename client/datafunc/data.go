package datafunc

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Data() {
	const (
		host     = "localhost"
		port     = "5432"
		user     = "postgres"
		password = "password123"
		dbname   = "vulnerable_app"
	)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("데이터베이스 연결 실패: %v", err)
		log.Println("로컬 환경에서 실행 중인지 확인하세요.")
		return
	}

	// 데이터베이스 연결 확인
	err = DB.Ping()
	if err != nil {
		log.Printf("데이터베이스 연결 확인 실패: %v", err)
		log.Println("PostgreSQL이 실행 중인지 확인하세요.")
		return
	}

	log.Println("데이터베이스 연결 성공")
}
