FROM golang:1-alpine as build

WORKDIR /build

COPY go.mod go.sum /build/

RUN go mod download

COPY . /build

RUN CGO_ENABLED=0 go build -o adguard-exporter .

FROM gcr.io/distroless/static-debian12

COPY --from=build /build/adguard-exporter /usr/bin/adguard-exporter

EXPOSE 9162

CMD ["/usr/bin/adguard-exporter"]

WORKDIR /

USER 1000
