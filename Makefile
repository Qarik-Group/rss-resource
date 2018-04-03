default: assets/check assets/in
assets/in:
	CGO_ENABLED=0 go build -o $@ ./cmd/in
assets/check:
	CGO_ENABLED=0 go build -o $@ ./cmd/check

docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o assets/check ./cmd/check
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o assets/in ./cmd/in
	docker build -t starkandwayne/rss-resource:latest .

all: default docker
.PHONY: assets/in assets/check docker all
