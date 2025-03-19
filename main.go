package main

import (
	"fmt"
	"hack/datafunc"
	"hack/service"
	"net/http"
)

func main() {
	datafunc.Data()

	service.SetDB(datafunc.DB)

	// 정적 파일 서빙
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// 라우트 설정
	http.HandleFunc("/", service.HandleLogin)
	http.HandleFunc("/login", service.HandleLoginPost)
	http.HandleFunc("/logout", service.HandleLogout)
	http.HandleFunc("/dashboard", service.HandleDashboard)
	http.HandleFunc("/post/create", service.HandleCreate)
	http.HandleFunc("/post", service.HandlePost)
	http.HandleFunc("/post/edit", service.HandleEditPost)
	http.HandleFunc("/post/update", service.HandleUpdatePost)
	http.HandleFunc("/post/delete", service.HandleDeletePost)
	// 서버 시작
	fmt.Println("서버가 시작되었습니다. http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
