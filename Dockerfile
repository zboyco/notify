# build
FROM golang:alpine AS builder
WORKDIR /usr/src/app
COPY ./ ./
RUN cd notify && go build -v .

# runtime
FROM scratch
COPY --from=builder /usr/src/app/notify/notify /go/bin/notify
EXPOSE 80
WORKDIR /go/bin
ENTRYPOINT ["/go/bin/notify"]