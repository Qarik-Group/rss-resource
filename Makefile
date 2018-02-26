default: assets/check assets/in
assets/in:
	go build -o $@ ./cmd/in
assets/check: check
	cp $+ $@
	chmod 755 $@

.PHONY: assets/in
