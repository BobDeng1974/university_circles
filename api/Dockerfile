FROM alpine:3.2
ADD api /api
COPY . .
ADD resource/pri_key.pem resource/pub_key.pem /resource/
#ADD resource/ca-certificates.crt /etc/ssl/certs/
VOLUME /resource/
ENTRYPOINT [ "/api" ]