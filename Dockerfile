FROM golang:1.20.0 as builder
MAINTAINER lanyulei <fdevops@163.com>

WORKDIR /opt/openiam
COPY . .

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy && make build

FROM alpine:3.17.1 as worker

WORKDIR /opt/openiam
COPY --from=builder /opt/openiam/openiam /opt/openiam/openiam
COPY --from=builder /opt/openiam/config/settings.yml /opt/openiam/config

EXPOSE 8000

ENTRYPOINT ["/opt/openiam/openiam", "server", "-c", "/opt/openiam/config/settings.yml"]
