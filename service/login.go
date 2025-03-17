package service

import (
	"database/sql"
	"net/http"
	"strconv"
)

var db *sql.DB

type Login interface {
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleDashboard(w http.ResponseWriter, r *http.Request)
	HandleLoginPost(w http.ResponseWriter, r *http.Request)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func HandleLoginPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// 취약한 SQL 쿼리 - 인젝션 가능
	query := "SELECT id FROM users WHERE username = '" + username + "' AND password = '" + password + "'"

	var userId int
	err := db.QueryRow(query).Scan(&userId)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// 쿠키 설정
	http.SetCookie(w, &http.Cookie{
		Name:  "user_id",
		Value: strconv.Itoa(userId),
		Path:  "/",
	})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
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
