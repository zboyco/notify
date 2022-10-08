# build
FROM golang:alpine AS builder
WORKDIR /usr/src/app
COPY ./ ./
RUN cd notify && go build -v .

# runtime
# FROM gcr.io/distroless/base:debug
FROM alpine
COPY --from=builder /usr/src/app/notify/notify /app
EXPOSE 80
CMD ["/app"]