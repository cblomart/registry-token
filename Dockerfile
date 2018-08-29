ARG TARGET=amd64
FROM alpine as builde
# Create nonroot user
RUN adduser -D -g '' registry-token-ldap
# Add ca-certificates
RUN apk --update add ca-certificates

FROM scratch
LABEL maintainer="cblomart@gmail.com"

# copy password file for users
COPY --from=builder /etc/passwd /etc/passwd

# copy ca-certificates 
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# copy binary
COPY ./releases/${TARGET}/registry-token-ldap /registry-token-ldap

#copy config
COPY ./registry-token-ldap.yml /etc/registry-token-ldap.yml

USER registry-token-ldap

VOLUME [ "/etc" ]

ENTRYPOINT ["/bin/registry-token-ldap", "-logtostderr"] 

