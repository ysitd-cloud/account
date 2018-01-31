FROM ysitd/glide as builder

COPY . /go/src/code.ysitd.cloud/component/account

WORKDIR /go/src/code.ysitd.cloud/component/account

RUN glide --no-color install -v --skip-test && \
    go build -v -ldflags="-s -w"

FROM alpine:3.6
COPY --from=builder /go/src/code.ysitd.cloud/component/account/account /

CMD ["/account"]
