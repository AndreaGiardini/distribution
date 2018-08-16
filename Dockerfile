FROM golang:1.10-alpine

ENV DISTRIBUTION_DIR /go/src/github.com/docker/distribution
ENV DOCKER_BUILDTAGS include_oss include_gcs

RUN set -ex \
    && apk add --no-cache make git gcc

WORKDIR $DISTRIBUTION_DIR
RUN go get github.com/anacrolix/torrent github.com/anacrolix/torrent/bencode github.com/anacrolix/torrent/metainfo
COPY . $DISTRIBUTION_DIR
COPY cmd/registry/config-example.yml /etc/docker/registry/config.yml

RUN make PREFIX=/go clean binaries

VOLUME ["/var/lib/registry"]
EXPOSE 5000
ENTRYPOINT ["registry"]
CMD ["serve", "/etc/docker/registry/config.yml"]
