FROM postgres:latest
ENV POSTGRES_PASSWORD=123
ENV POSTGRES_DB=test
ENV POSTGRES_USER=postgres
EXPOSE 5432