# build
FROM golang:alpine AS builder
WORKDIR /usr/src/app
COPY ./ ./
RUN cd notify && CGO_ENABLED=0 go build -v -ldflags '-w -s' .

# upx
FROM ghcr.io/zboyco/upx:alpine AS upx
COPY --from=builder /usr/src/app/notify/notify /app
RUN upx --best /app

# runtime
FROM gcr.io/distroless/static
COPY --from=upx /app /app
EXPOSE 80
ENTRYPOINT ["/app"]