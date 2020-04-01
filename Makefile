GO_TEST_FLAGS?=-race -short -covermode=atomic -coverprofile=coverage.out
GO_TEST_PKGS?=$(shell go list ./...)

.PHONY: test
test:
	go test ${GO_TEST_FLAGS} ${GO_TEST_PKGS}

.PHONY: fmt
fmt:
	@go fmt  ./...

.PHONY: fmt-check
fmt-check:
	# get all go files and run go fmt on them
	@diff=$$(go fmt  ./...); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;
