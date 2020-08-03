FROM golang as build

ADD . /usr/local/go/src/mili
WORKDIR /usr/local/go/src/mili

RUN CGO_ENBLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.9
ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN=""
ENV GIN_MODE="release"
ENV PORT=3000

# 设置alpine的镜像地址为阿里云的地址
RUN echo "https://mirrors.aliyun.com/alpine/v3.9/main/" > /etc/apk/repositories \
    # 安装依赖包
    && apk update \
    apk add ca-certificates && \
    echo "hosts: file dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /usr/local/go/src/mili/api_server /usr/bin/api_server
ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]