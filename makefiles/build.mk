SERVICE_NAME=blockchain-api
GIT_COMMIT=$(shell git log -1 --pretty=format:"%H")


.PHONY: go-build
build-binary: ## Build service binary
	@echo "==> Building go binary"
	@env GOOS=linux go build -ldflags="-s -w" -o ./build/artifacts/$(SERVICE_NAME) golang/cmd/main.go
	@echo "==> Done. Check './build/artifacts' dir"


.PHONY: build-docker-image
build-docker-image: graphql build-binary
	@echo "==> Building docker image"
	docker build \
		-t $(SERVICE_NAME) \
		--build-arg GIT_COMMIT="$(GIT_COMMIT)" \
		-f ./build/Dockerfile .


