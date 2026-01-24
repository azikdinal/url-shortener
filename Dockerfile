FROM golang:1.25.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/  ./
COPY internal/  ./

EXPOSE 8000

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-url-shortener

CMD ["/docker-url-shortener"]
