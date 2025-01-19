.DEFAULT_GOAL := test
.SILENT:

.PHONY: test
test:
	go test -v ./...
