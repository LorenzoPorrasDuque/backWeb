FROM golang:alpine3.19
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV   DB_USER=postgres
ENV   DB_PASSWORD=123
ENV   DB_NAME=test
ENTRYPOINT ["go", "run", "main.go"]
