FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /mili_server
WORKDIR /mili_server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.9

ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN=""
ENV GIN_MODE="release"
ENV PORT=3000

# 设置alpine的镜像地址为阿里云的地址
RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /mili_server/api_server /usr/bin/api_server
ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]