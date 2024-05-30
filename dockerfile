FROM alpine:3.19


RUN mkdir /go && cd /go \
    && wget --no-check-certificate https://golang.google.cn/dl/go1.20.linux-amd64.tar.gz \
    && tar -C /usr/local -zxf go1.20.linux-amd64.tar.gz \
    && rm -rf /go/go1.20.linux-amd64.tar.gz 

RUN apk add gcompat

ENV GOPATH=/go
ENV PATH=/usr/local/go/bin:$GOPATH/bin:$PATH
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /home

COPY . .

# CMD ["go","run","./cmd/main.go"]