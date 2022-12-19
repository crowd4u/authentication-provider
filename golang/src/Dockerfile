ARG SHA256_KEY="hogehoge"
FROM golang:latest as builder
WORKDIR /workdir
COPY . .
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
# Build
RUN go build -o app

# Runtime Container
FROM gcr.io/distroless/base
COPY --from=builder /workdir/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
