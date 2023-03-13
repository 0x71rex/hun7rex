FROM golang:1.20.2-alpine as build-env
RUN GO111MODULE=on go get -v github.com/0x71rex/hun7rex/cmd/hun7rex

FROM alpine:latest
RUN apk add --no-cache bind-tools ca-certificates
COPY --from=build-env /go/bin/hun7rex /usr/local/bin/hun7rex
ENTRYPOINT ["hun7rex"]