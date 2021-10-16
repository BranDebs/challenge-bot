FROM golang:1.17-alpine AS builder
# git for fetching dependencies.
# ca-certificate for calling HTTPS endpoints.
RUN apk update && \
    apk add --no-cache git ca-certificates wget xz && \
    update-ca-certificates
RUN wget https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz -P /usr/local
RUN xz -d -c /usr/local/upx-3.96-amd64_linux.tar.xz | tar -xOf - upx-3.96-amd64_linux/upx > /bin/upx && chmod a+x /bin/upx

# Create unprivilege user.
ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main
RUN upx -9 main

FROM alpine

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY config.toml /app/config.toml
COPY apitoken /app/apitoken

# Copy static executable.
COPY --from=builder /app/main /app/main

# Use the unprivileged user.
USER appuser:appuser

ENTRYPOINT [ "/app/main" ]
