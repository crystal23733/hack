package service

import (
	"database/sql"
	"net/http"
	"strconv"
)

var db *sql.DB

// 데이터베이스 연결 설정
func SetDB(database *sql.DB) {
	db = database
}

type Login interface {
	HandleLogin(w http.ResponseWriter, r *http.Request)
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
