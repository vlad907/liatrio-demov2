# builder stage: compile the Go app with the same toolchain we declare in go.mod
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
# - go.mod names the module, pins dependency versions, and keeps builds reproducible (think requirements.txt for Go)
COPY go.sum ./
# - go.sum carries the checksums for those modules so the download can be verified and tamper-free
# pull module dependencies up front so later source changes can reuse cache
RUN go mod download
# - fetch all dependencies ahead of time so we reuse this layer if only code changes
COPY . .
# - bring the rest of the application source into /app

# build a static linux/amd64 binary that Cloud Run can run without extra libs
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server
# - compile a static Linux/amd64 binary named server so Cloud Run can execute it directly

# runtime stage: distroless keeps the image small and secure (no shell, no package manager)
FROM gcr.io/distroless/base-debian12
# - ultra-minimal Debian 12 base; no shell or package manager so the attack surface stays tiny
WORKDIR /app

# copy the compiled binary from the builder image into the minimal runtime
COPY --from=builder /app/server .

# match our workflow expectations by listening on port 80 inside the container
ENV PORT=80

# advertise that this container serves traffic on port 80
EXPOSE 80

ENTRYPOINT ["./server"]
# - When we do "docker run" it runs this binary