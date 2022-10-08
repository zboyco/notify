# build
FROM golang:alpine AS builder
WORKDIR /usr/src/app
COPY ./ ./
RUN cd notify && go build -v .

# runtime
FROM scratch
COPY --from=builder /usr/src/app/notify/notify /app
EXPOSE 80
CMD ["/app"]