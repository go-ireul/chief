FROM alpine:3.6

RUN mkdir -p /var/lib/chief

WORKDIR /var/lib/chief

VOLUME /var/lib/chief

ADD chief /bin

EXPOSE 9000

ENTRYPOINT ["/bin/chief"]
