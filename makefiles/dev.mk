.PHONY: graphql
graphql:  ## Generate GraphQL resolver from schema
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen


.PHONY: dev-service-start
dev-service-start:
	go run golang/cmd/main.go


.PHONY: run-docker-container
run-docker-container: build-docker-image
	docker run --name blockchain-api -p 8080:8080 --env-file .env.local -d --rm blockchain-api


.PHONY: stop-docker-container
stop-docker-container:
	docker stop blockchain-api









