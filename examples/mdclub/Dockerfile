FROM alpine
ADD mdclub.tar.gz /app/
RUN apk -Uuv add --no-cache ca-certificates tini tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
WORKDIR /app
ENTRYPOINT ["/sbin/tini","--", "mdclub"]