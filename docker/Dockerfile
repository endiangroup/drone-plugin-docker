# docker build --rm -f docker/Dockerfile -t plugins/docker .

FROM docker:17.12.0-ce-dind

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

ADD release/linux/amd64/drone-docker /bin/
ENTRYPOINT PLUGIN_TAGS=latest,`git describe --tags` /usr/local/bin/dockerd-entrypoint.sh /bin/drone-docker
