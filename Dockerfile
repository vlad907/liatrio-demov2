FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
# 👇 Ensure it's executable
RUN chmod +x /app/app
# 👇 Ensure root can bind to port 80
USER root
EXPOSE 80
CMD ["./app"]