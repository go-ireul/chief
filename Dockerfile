FROM scratch

VOLUME /var/lib/chief

WORKDIR /var/lib/chief

ADD chief /

EXPOSE 9000

ENTRYPOINT ["/chief"]
