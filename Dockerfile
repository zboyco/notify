# build
FROM golang:alpine AS builder
WORKDIR /usr/src/app
COPY ./ ./
RUN cd notify && CGO_ENABLED=0 go build -v .

# runtime
FROM gcr.io/distroless/static
COPY --from=builder /usr/src/app/notify/notify /app
EXPOSE 80
ENTRYPOINT ["/app"]