FROM alpine as builder

# Create nonroot user
RUN adduser -D -g '' registry-token-ldap

# Add ca-certificates
RUN apk --update add ca-certificates

FROM scratch
LABEL maintainer="cblomart@gmail.com"
ARG release_type=amd64

# copy password file for users
COPY --from=builder /etc/passwd /etc/passwd

# copy ca-certificates 
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# copy config
COPY ./config.yml /etc/registry-token-ldap/config.yml

# copy binary
COPY ./releases/${release_type}/registry-token-ldap /registry-token-ldap

# run as user 
USER registry-token-ldap

# allow sharing of /etc
VOLUME [ "/etc" ]

# start with logging
ENTRYPOINT ["/bin/registry-token-ldap", "-logtostderr"] 