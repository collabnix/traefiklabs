IMAGE ?= davidcollom/rancher-traefik
TAG ?= latest

build:
	docker build -t ${IMAGE}:${TAG} .

push: build
	docker push ${IMAGE}:${TAG}

run: build
	docker run -ti  -v $(shell pwd)/test_secrets/:/run/secrets/ --env-file=.env ${IMAGE}:${TAG}
