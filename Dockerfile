FROM --platform=linux/arm64,linux/amd64 golang:1-alpine

WORKDIR /usr/src/app

COPY . .

RUN go build -o /usr/bin/adguard-exporter && \
    rm -rf /usr/src/app

EXPOSE 9162

CMD ["/usr/bin/adguard-exporter"]

WORKDIR /

USER 1000
