FROM ysitd/glide as builder

COPY . /go/src/github.com/ysitd-cloud/account

WORKDIR /go/src/github.com/ysitd-cloud/account

RUN glide --no-color install -v && \
    go build -v

FROM alpine:3.6
COPY --from=builder /go/src/github.com/ysitd-cloud/account/account /

CMD ["/account"]
