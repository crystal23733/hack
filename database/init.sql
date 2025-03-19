-- 기존 데이터베이스가 있다면 삭제
DROP DATABASE IF EXISTS vulnerable_app;

-- 데이터베이스 생성
CREATE DATABASE vulnerable_app;

\c vulnerable_app

-- 사용자 테이블 생성
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL
);

-- 게시물 테이블 생성
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 기본 사용자 생성 (username: admin, password: password123)
INSERT INTO users (username, password) VALUES ('admin', 'password123');