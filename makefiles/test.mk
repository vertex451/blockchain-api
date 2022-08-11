# Unit test section
.PHONY: test-unit
test-unit: ## Unit testing
	go test -v ./golang/internal/...

.PHONY: mocks
mocks: clean-mocks ## For generating mock based on all project interfaces
	mockery --all --dir "./golang" --inpackage --case underscore


.PHONY: clean-mocks
clean-mocks: ## Cleans old mocks
	find . -name "mock_*.go" -type f -print0 | xargs -0 /bin/rm -f


# Integration test section
.PHONY: test-integration
test-integration: run-docker-container
	@echo "==> Running integration tests"
	go test -v ./golang/tests/integration/...
	make stop-docker-container

