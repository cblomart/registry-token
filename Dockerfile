FROM golang:latest AS build
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o registry-token-ldap .

FROM busybox
COPY -from build /app/registry-token-ldap /bin/
CMD ["registry-token-ldap -logtostderr"]