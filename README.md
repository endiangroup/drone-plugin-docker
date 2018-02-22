# drone-docker

Drone plugin to build and publish Docker images to a container registry.

This fork adds the facility to tag images using the current git tag. The functionality `auto_tag` has been replaced, and can be used as follows in a .drone.yml:

```
  some_pipeline:
    image: kowalatech/drone-docker
    repo: kowalatech/faucet
    secrets: [ docker_username, docker_password ]
    tag: my-tag
    privileged: true # may be required for docker-in-docker
    auto_tag: true   # adds the durrent git tag
    dockerfile: faucet.Dockerfile
    when:
      event: [push]
      branch: [master]
```

In that pipeline, git pushes to the master branch would result in a docker image tagged with `my-tag` and the current git tag, if ther is one.

## Build

Build the binary with the following commands:

```
sh .drone.sh
```

## Docker

Build the Docker image with the following commands:

```
docker build --rm=true -f docker/Dockerfile -t plugins/docker .
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e PLUGIN_TAG=latest \
  -e PLUGIN_REPO=octocat/hello-world \
  -e DRONE_COMMIT_SHA=d8dbe4d94f15fe89232e0402c6e8a0ddf21af3ab \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  --privileged \
  plugins/docker --dry-run
```
