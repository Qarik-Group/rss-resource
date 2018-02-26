default: assets/check assets/in
assets/in:
	CGO_ENABLED=0 go build -o $@ ./cmd/in
assets/check:
	CGO_ENABLED=0 go build -o $@ ./cmd/check

docker:
	docker build -t starkandwayne/rss-resource:latest .

all: default docker
.PHONY: assets/in assets/check docker all
