# SPDX-License-Identifier: ISC
# Copyright (c) 2019-2021 Bitmark Inc.
# Use of this source code is governed by an ISC
# license that can be found in the LICENSE file.

.PHONY:

dist =
GITHUB_USER =
GITHUB_TOKEN =

ARCH = $(shell /usr/bin/uname -m)

DOCKER_BUILD_COMMAND = docker build

ifeq ($(ARCH), arm64)
DOCKER_BUILD_COMMAND = docker buildx build --platform linux/amd64 --load
endif

default: build

config:

logistic-server:
	go build -o bin/logistic-server ./services/logistic-server

run-logistic-server: logistic-server
	./bin/logistic-server

build: logistic-server

run: run-logistic-server

build-logistic-server:
ifndef dist
	$(error dist is undefined)
endif
	$(DOCKER_BUILD_COMMAND) --build-arg dist=$(dist) \
	--build-arg GITHUB_USER=$(GITHUB_USER) \
	--build-arg GITHUB_TOKEN=$(GITHUB_TOKEN) \
	-t logistic-server:api-$(dist) .
	docker tag logistic-server:api-$(dist) 083397868157.dkr.ecr.ap-northeast-1.amazonaws.com/logistic-server:api-$(dist)

image: build-logistic-server

push:
ifndef dist
	$(error dist is undefined)
endif
	aws ecr get-login-password | docker login --username AWS --password-stdin 083397868157.dkr.ecr.ap-northeast-1.amazonaws.com
	docker push 083397868157.dkr.ecr.ap-northeast-1.amazonaws.com/logistics-tool:$(dist)

test:
	go test ./...

clean:
	rm -rf bin
