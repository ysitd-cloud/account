FROM ysitd/dep AS builder

WORKDIR /go/src/code.ysitd.cloud/component/account

COPY . /go/src/code.ysitd.cloud/component/account

RUN dep ensure -vendor-only && \
    go build -o account code.ysitd.cloud/component/account/cmd

FROM alpine:3.6

COPY --from=builder /go/bin/code.ysitd.cloud/component/account/server /

CMD ["/account"]