REPO:=ttl.sh/knabben/dos-poc
TAG:=latest

.PHONY: build
build:
	docker build -t ${REPO}:${TAG} -f Dockerfile .

.PHONY: test
test:
	go test -v ./...
