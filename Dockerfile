FROM golang:alpine AS build
RUN apk update && apk add ca-certificates && apk add git && apk add upx && rm -rf /var/cache/apk/*
RUN mkdir -p /go/src/github.com/cblomart/registry-token-ldap
ADD . /go/src/github.com/cblomart/registry-token-ldap/
WORKDIR /go/src/github.com/cblomart/registry-token-ldap/
RUN go get ./...
RUN go build  -ldflags="-s -w" -o registry-token-ldap . && upx -9 -q registry-token-ldap

FROM alpine
COPY --from=build /etc/ssl/certs /etc/ssl/certs
COPY --from=build /go/src/github.com/cblomart/registry-token-ldap/registry-token-ldap /bin/
RUN mkdir /etc/registry-token-ldap/
ADD config.yml /etc/registry-token-ldap/
CMD ["/bin/registry-token-ldap", "-logtostderr"]