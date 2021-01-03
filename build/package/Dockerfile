FROM alpine:latest


ENV LANG="zh_CN.UTF-8"

RUN apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

ADD ./agent /home/agent

WORKDIR /home/agent

RUN chmod a+x caesar

ENTRYPOINT exec ./caesar "$0" "$@"