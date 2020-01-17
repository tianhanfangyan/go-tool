
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

.PHONY: vendor
vendor:
	go mod vendor

fmt:
	@gofmt -w ${GOFILES}
