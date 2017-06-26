FROM golang:1.8-alpine

ADD . /go/src/github.com/ysitd-cloud/account
WORKDIR /go/src/github.com/ysitd-cloud/account

RUN go install

EXPOSE 8080

CMD ["account"]
