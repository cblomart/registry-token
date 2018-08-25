FROM golang:alpine AS build
RUN apk update && apk add ca-certificates && apk add git && rm -rf /var/cache/apk/*
RUN mkdir -p /go/src/github.com/cblomart/registry-token-ldap
ADD . /go/src/github.com/cblomart/registry-token-ldap/
WORKDIR /go/src/github.com/cblomart/registry-token-ldap/
RUN go get ./...
RUN go build -o registry-token-ldap .

FROM alpine
COPY --from=build /etc/ssl/certs /etc/ssl/certs
COPY --from=build /go/src/github.com/cblomart/registry-token-ldap/registry-token-ldap /bin/
ADD registry-token-ldap.yml /etc/
CMD ["/bin/registry-token-ldap", "-logtostderr"]