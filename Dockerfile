FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 80
CMD ["./app"]