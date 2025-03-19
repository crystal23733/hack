package service

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Dashboard interface {
	HandleDashboard(w http.ResponseWriter, r *http.Request)
	HandleCreate(w http.ResponseWriter, r *http.Request)
	HandlePost(w http.ResponseWriter, r *http.Request)
	HandleEditPost(w http.ResponseWriter, r *http.Request)
	HandleUpdatePost(w http.ResponseWriter, r *http.Request)
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

func HandlePost(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, content, created_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var posts []map[string]interface{}
	for rows.Next() {
		var id int
		var title, content string
		var createdAt string
		if err := rows.Scan(&id, &title, &content, &createdAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		post := map[string]interface{}{
			"id":         id,
			"title":      title,
			"content":    content,
			"created_at": createdAt,
		}
		posts = append(posts, post)
	}

	// HTML 직접 생성 (XSS 취약)
	w.Header().Set("Content-Type", "text/html")
	for _, post := range posts {
		// 이스케이프 X
		fmt.Fprintf(w, `
        <div class="post" data-id="%d">
            <h3>%s</h3>
            <p>%s</p>
            <small>작성일: %s</small>
            <div>
                <button onclick="editPost(%d)">수정</button>
                <button onclick="deletePost(%d)">삭제</button>
            </div>
        </div>
        `, post["id"], post["title"], post["content"], post["created_at"], post["id"], post["id"])
	}
}

func HandleEditPost(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("id")

	query := "SELECT title, content FROM posts WHERE id = " + postID

	var title, content string
	err := db.QueryRow(query).Scan(&title, &content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl := template.New("edit")
	tmpl, _ = tmpl.Parse(`
	<form id="edit-form" method="POST" action="/post/update?id={{ .ID }}">
		<input type="hidden" name="id" value="{{.ID}}">
		<div>
			<label for="title">제목:</label>
			<input type="text" id="title" name="title" value="{{.Title}}">
		</div>
		<div>
			<label for="content">내용:</label>
			<textarea id="content" name="content">{{.Content}}</textarea>
		</div>
		<button type="submit">수정</button>
		<button type="button" onclick="window.location='/dashboard'">취소</button>
	</form>
	`)

	tmpl.Execute(w, map[string]interface{}{
		"ID":      template.HTML(postID),
		"Title":   template.HTML(title),
		"Content": template.HTML(content),
	})
}

func HandleUpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	postID := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	query := fmt.Sprintf("UPDATE posts SET title = '%s', content = '%s' WHERE id = %s", title, content, postID)

	_, err := db.Exec(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
