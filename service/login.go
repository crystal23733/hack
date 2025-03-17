package service

import "net/http"

type Login interface {
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleDashboard(w http.ResponseWriter, r *http.Request)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/dashboard.html")
}
