FROM ysitd/dep as builder

COPY . /go/src/code.ysitd.cloud/auth/account

WORKDIR /go/src/code.ysitd.cloud/auth/account

RUN dep ensure && \
    go build -v -ldflags="-s -w"

FROM alpine:3.6
COPY --from=builder /go/src/code.ysitd.cloud/component/account/account /

CMD ["/account"]
