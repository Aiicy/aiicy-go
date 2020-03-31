.PHONY: test
test:
	@go test ${GO_TEST_FLAGS} ${GO_TEST_PKGS}
	@go tool cover -func=coverage.out | grep total

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
