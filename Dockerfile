FROM alpine:3.4

ADD account /
ADD dist /

CMD ["/account"]
