FROM golang:1.16.5-buster as builder
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . . 
RUN GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o flusher main.go
FROM busybox as downloader
RUN wget -O /tmp/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64
RUN chmod +x /tmp/dumb-init

FROM bitnami/minideb:buster
ENV TZ="Asia/Jakarta" FLUSHER_VERSION="0.7.6"
COPY --from=builder /app/flusher /app/flusher
COPY --from=downloader /tmp/dumb-init /usr/bin/dumb-init
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/app/flusher"]