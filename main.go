package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 정적 파일 서빙
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// 라우트 설정
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/dashboard.html")
	})

	// 서버 시작
	fmt.Println("서버가 시작되었습니다. http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
