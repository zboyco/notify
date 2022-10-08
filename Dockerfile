# build
FROM golang:alpine AS builder
WORKDIR /usr/src/app
COPY ./ ./
RUN cd notify && go build -v .

# runtime
FROM alpine
COPY --from=builder /usr/src/app/notify/notify /go/bin/notify
EXPOSE 80
WORKDIR /go/bin
ENTRYPOINT ["./notify"]