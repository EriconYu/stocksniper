FROM golang:1.8.3
MAINTAINER Ericon Yu <yuzhiqiang@stackbang.com>
RUN cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY ./src/. $GOPATH/src/
COPY ./bin/. $GOPATH/bin/
WORKDIR $GOPATH/bin

RUN go build stocksniper

EXPOSE 8080
CMD ["$GOPATH/bin/stocksniper"]
