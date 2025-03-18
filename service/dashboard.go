package service

import (
	"fmt"
	"net/http"
	"strconv"
)

type Dashboard interface {
	HandleDashboard(w http.ResponseWriter, r *http.Request)
	HandleCreate(w http.ResponseWriter, r *http.Request)
}

func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	// 인증 확인
	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// 사용자 ID가 유효한지 확인 (SQL 인젝션에 취약한 방식으로 구현)
	userId := cookie.Value
	query := "SELECT id FROM users WHERE id = " + userId

	var id int
	err = db.QueryRow(query).Scan(&id)
	if err != nil {
		// 유효하지 않은 사용자 ID인 경우 로그인 페이지로 리다이렉트
		http.SetCookie(w, &http.Cookie{
			Name:   "user_id",
			Value:  "",
			MaxAge: -1,
			Path:   "/",
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "./static/dashboard.html")
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusUnauthorized)
		return
	}

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, "로그인이 필요합니다", http.StatusUnauthorized)
		return
	}

	userId, _ := strconv.Atoi(cookie.Value)
	title := r.FormValue("title")
	content := r.FormValue("content")

	// 취약한 SQL 쿼리 - 인젝션 가능
	query := fmt.Sprintf("INSERT INTO posts (user_id, title, content) VALUES (%d, '%s', '%s')", userId, title, content)

	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
