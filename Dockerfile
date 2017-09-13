FROM ysitd/glide

COPY . /go/src/github.com/ysitd-cloud/account

WORKDIR /go/src/github.com/ysitd-cloud/account

RUN glide --no-color install -v && \
    go install -v

ENV PORT 80
EXPOSE 80

CMD ["account"]
