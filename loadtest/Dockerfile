FROM alpine:3.10.0
LABEL author="william.pjyeh@gmail.com"

RUN apk add --no-cache bash curl

COPY loadtest.sh /
RUN  chmod +x /loadtest.sh

ENTRYPOINT ["/loadtest.sh"]
