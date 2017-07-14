FROM golang:1.8-alpine

COPY . /go/src/github.com/ysitd-cloud/account

WORKDIR /go/src/github.com/ysitd-cloud/account

RUN wget -qO- https://glide.sh/get | sh && \
    glide install && \
    go install

CMD ["account"]
