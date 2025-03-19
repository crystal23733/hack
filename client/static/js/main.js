// 페이지 로드 시 실행
document.addEventListener("DOMContentLoaded", () => {
    const loginForm = document.getElementById("login-form");
    loginForm.addEventListener("submit", (e) =>{
        e.preventDefault();

        const username = document.getElementById("username").value;
        const password = document.getElementById("password").value;

        if(!username || !password){
            alert("사용자명과 비밀번호를 입력해주세요.");
            return;   
        }

        // 폼 데이터 준비
        const formData = new FormData();
        formData.append("username", username);
        formData.append("password", password);

        // 로그인 요청 보내기
        fetch("/login", {
          method:"POST",
          body:formData  
        })
        .then(response => {
            if(response.redirected){
                // 성공적으로 로그인되면 리다이렉트된 URL로 이동
                window.location.href = response.url;
            } else if(response.status === 401) {
                // 인증 실패 시
                return response.text().then(text => {
                    alert("로그인 실패: " + text);
                });
            } else {
                // 다른 오류
                alert("서버 오류");
            }
        })
        .catch(error => {
            console.error("로그인 오류:", error);
            alert("로그인 오류가 발생했습니다. 다시 시도해주세요.");
        });
    });
});