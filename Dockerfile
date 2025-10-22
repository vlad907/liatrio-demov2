FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
RUN chmod +x /app/app
USER root
EXPOSE 80
CMD ["./app"]