# Use the builder image
FROM golang:1.18 as builder
WORKDIR /app
COPY . .
RUN echo "[+] Downloading dependencies."
RUN go mod download

RUN echo "[+] Starting intermediate build "
RUN start=$(date +%s) && \
    CGO_ENABLED=0 GOOS=linux go build -o app &&\
    end=$(date +%s) && \
    echo "[+] intermediate Build completed in $((end-start)) seconds"

# Create a directory for the config volume
RUN mkdir /config

# multistage build
FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN echo "[+] Copying binary from intermediate build to final build."
COPY --from=builder /app/app /app

# Define a volume for the config directory
RUN echo "[+] Attaching config volume"
VOLUME ["/config"]

EXPOSE 8080

RUN echo "[+] Running binary"
RUN chmod +x ./app
CMD ["./app"]