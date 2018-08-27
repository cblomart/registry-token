FROM golang:alpine AS builder
RUN apk update && apk add ca-certificates && apk add git && apk add upx && rm -rf /var/cache/apk/*
COPY . /go/src/github.com/cblomart/registry-token-ldap/
WORKDIR /go/src/github.com/cblomart/registry-token-ldap/
RUN go get ./...
RUN go build  -ldflags="-s -w" -o registry-token-ldap . && upx -9 -q registry-token-ldap

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/cblomart/registry-token-ldap/registry-token-ldap /bin/
COPY config.yml /etc/registry-token-ldap/
CMD ["/bin/registry-token-ldap", "-logtostderr"]