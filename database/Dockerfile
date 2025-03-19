FROM postgres:15

# 환경 변수 설정
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password123

# 초기화 스크립트 복사
COPY init.sql /docker-entrypoint-initdb.d/
COPY pg_hba.conf /etc/postgresql/pg_hba.conf

# 권한 설정
USER root
RUN chown postgres:postgres /etc/postgresql/pg_hba.conf

# 포트 설정
EXPOSE 5432

