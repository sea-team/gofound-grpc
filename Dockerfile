FROM golang:1.18 as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /go/workdir

ADD ../.. .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o gofound ./
RUN go build -ldflags="-s -w" -installsuffix cgo -o gateway ./gateway

FROM bash:latest as prod

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/workdir/gofound /
COPY --from=build /go/workdir/gateway /
COPY --from=build /go/workdir/config/config.yaml /
COPY --from=build /go/workdir/start.sh /
RUN chmod +x start.sh

ENTRYPOINT ["sh", "start.sh"]
EXPOSE 5678
EXPOSE 4567



