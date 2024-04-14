FROM golang:alpine3.19
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
ENTRYPOINT ["go", "run", "main.go"]
