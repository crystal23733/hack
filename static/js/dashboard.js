document.addEventListener("DOMContentLoaded", () => {
    checkLoginStatus();
    loadPosts();
    logout();
    const postForm = document.getElementById("post-form");
    postForm.addEventListener("submit", (e) => {
        e.preventDefault();

        const title = document.getElementById("title").value;
        const content = document.getElementById("content").value;

        if (!title || !content) {
            alert("제목과 내용을 입력해주세요.");
            return;
        }

        const formData = new FormData();
        formData.append("title", title);
        formData.appent("content", content);

        // 게시글 생성 요청
        fetch("/dashboard/create", {
            method: "POST",
            body: formData,
        })
        .then(response => {
            if (response.redirected) {
                window.location.href = response.url;
            } else if (response.status === 401) {
                alert("로그인이 필요합니다.");
                window.location.href = "/";
            } else {
                alert("게시글 생성 실패");
            }
        })
        .catch(error => {
            console.error("게시글 생성 오류:", error);
            alert("게시글 생성 오류가 발생했습니다. 다시 시도해주세요.");
        });
    })
})

const loadPosts = () => {
    fetch("/post")
    .then(response => response.text())
    .then(html => {
        document.getElementById("posts-container").innerHTML = html;
    })
    .catch(error => {
        console.error("게시물을 가져오는 중 오류 발생:", error);
    })
}

const checkLoginStatus = () => {
    const cookies = document.cookie.split(";");
    let userIdExists = false;

    for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i].trim();
        if (cookie.startsWith("user_id=") && cookie.substring(8) !== '') {
            userIdExists = true;
            break;
        }
    }

    if (!userIdExists) {
        // 쿠키가 없으면 로그인 페이지로 즉시 리다이렉트
        window.location.href = "/";
        return false;
    }
    return true;
}

const logout = () => {
    const logoutLink = document.querySelector('.logout');
    
    logoutLink.addEventListener('click', function(event) {
        event.preventDefault();
        
        fetch('/logout', {
            method: 'POST',
            credentials: 'same-origin'
        })
        .then(response => {
            document.cookie = "user_id=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
            window.location.href = response.url;
        })
        .catch(error => {
            console.error("로그아웃 오류:", error);
            alert("로그아웃 중 오류가 발생했습니다.");
        });
    });
}