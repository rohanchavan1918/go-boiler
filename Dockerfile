FROM golang:1.18 as builder
WORKDIR /app
COPY . .
RUN pwd
RUN ls
RUN ls ./config/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app
# multistage build
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app /app
EXPOSE 8080
CMD ["./app"]