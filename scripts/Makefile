DOCKER_ENV        ?= DOCKER_BUILDKIT=1
DOCKER_BUILD_ARGS ?= --rm
DOCKER_TAG        ?= 1.0
DOCKER_REGISTRY   ?= docker.io
DOCKER_REPOSITORY ?= monandkey/rcmd

rcmd              ?= ${DOCKER_REGISTRY}/${DOCKER_REPOSITORY}:${DOCKER_TAG}

build-all: build-rcmd copy-rcmd

.PHONY: build-rcmd
build-rcmd:
	${DOCKER_ENV} docker build ${DOCKER_BUILD_ARGS} \
		--tag ${rcmd} \
		--file ./build/Dockerfile \
		--no-cache \
		./

.PHONY: copy-rcmd
copy-rcmd:
	docker run -itd --name rcmd ${DOCKER_REPOSITORY}:${DOCKER_TAG} /bin/sh
	docker container cp rcmd:/home/rcmd-windows.zip ../bin/
	docker container stop rcmd
	docker container rm rcmd
