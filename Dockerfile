# build
FROM golang:alpine AS builder
WORKDIR /usr/src/app
COPY ./ ./
RUN cd notify && go build -v .

# runtime
FROM gcr.io/distroless/static-debian10:debug
COPY --from=builder /usr/src/app/notify/notify /app
EXPOSE 80
CMD ["/app"]