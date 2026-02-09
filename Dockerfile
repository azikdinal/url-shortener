# build stage
FROM golang:1.25.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/  ./cmd/
COPY internal/  ./internal/
COPY gen/  ./gen/
COPY api/  ./api/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -o app ./cmd/main/main.go

# runtime stage
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/app /app/app

EXPOSE 8000 8001
USER nonroot:nonroot
CMD ["/app/app"]
