FROM nginx:alpine

# nginx 설정 파일 복사
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]