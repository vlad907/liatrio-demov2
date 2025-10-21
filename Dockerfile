FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]