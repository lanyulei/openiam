FROM golang:1.20.0 as builder
MAINTAINER lanyulei <fdevops@163.com>

WORKDIR /opt/openops
COPY . .

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy && make build

FROM alpine:3.17.1 as worker

WORKDIR /opt/openops
COPY --from=builder /opt/openops/openops /opt/openops/openops
COPY --from=builder /opt/openops/config/settings.yaml /opt/openops/config

EXPOSE 8000

ENTRYPOINT ["/opt/openops/openops", "server", "-c", "/opt/openops/config/settings.yaml"]
