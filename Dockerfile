FROM golang:1.8.1

MAINTAINER ztime "ztime@wukc.com"  

RUN mkdir -p /go/src/work-codes/ && cd /go/src/work-codes

RUN git clone https://code.aliyun.com/wukc/wuyue.git

WORKDIR /go/src/work-codes/wuyue/app

RUN go get github.com/Masterminds/glide

RUN glide install 

RUN go build

EXPOSE 1323
CMD [ "./app" ]