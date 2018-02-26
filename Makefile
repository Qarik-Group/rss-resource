default: assets/check assets/in
assets/in:
	CGO_ENABLED=0 go build -o $@ ./cmd/in
assets/check: check
	cp $+ $@
	chmod 755 $@

docker:
	docker build -t starkandwayne/rss-resource .

all: default docker
.PHONY: assets/in docker all
