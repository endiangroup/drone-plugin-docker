#!/bin/sh

set -e
set -x

# compile the main binary
GOOS=linux GOARCH=arm   CGO_ENABLED=0 GOARM=7 go build -ldflags "-X main.build=${DRONE_BUILD_NUMBER}" -a -tags netgo -o release/linux/arm/drone-docker   github.com/drone-plugins/drone-docker/cmd/drone-docker
