# Blockchain API

## Required steps
### Prerequisites
- Check if you have golang version not higher than v1.17(mockery lib is not migrated yet)  (https://go.dev/doc/install)
- Go modules should be enabled. Set in your bash_profile env `GO111MODULE=on` 
and also check corresponding preference in your IDE. If you use Gomodules, you don't need to set up $GOROOT or $GOPATH/
- Enable `goimports` watcher in your IDE(it automatically formats code and deals with imports)

### Start service
- Install needed dependencies `go mod tidy`
- Generate GraphQL files  `make graphql`. Otherwise, you will see the `Module is not in GOROOT` error.
- Start the service by `make dev-service-start`. You should see `"connected to GraphQL playground"` message in log. 

### Before committing:
- Check linter by `make lint` (See Linter section for details)
- Check Unit tests by `make test-unit` (See Unit tests section for details)
- Check Integration tests by `make test-integration` (See Integration tests section for details)

## Project structure
```
internal: is for modules that is not likely to be used in other projects
└── service: wrapper around modules following Clean Architecture approach
    └── transport - transport layer
    └── usecase - business logic layer
    └── repo - repository layer
pkg: is for modules that can be used in other projects, in future we can move them to separate library
```

## Linter
To run linter run `make lint` command

We use golangci linter, that aggregates all available linters.
Linter's configuration is located in `.golangci.yml`. 
I started from strict linters, feel free to make it less strict.

## Unit tests
To run unit tests we need to generate mock stubs first by `make mocks` command.
For this you should have installed `https://github.com/vektra/mockery`.

This library gives us an ability to mock everything with defining the request params and return values.

## Integration tests
`make test-integration` will spin up the service in docker container and will query GraphQL endpoints.
So it covers the all layers from transport to the blockchain or database.
Also, it checks build, so if something goes wrong with it, integration tests will let us know.

## Build the binary
`make build-binary` will build the code into binary file and place in /build/artifacts directory. 
You can run it right away by just calling that file(`./build/artifacts/blockchain-api`)

## Run in Docker container
`run-docker-container` will build the binary, than docker image and will run the docker container with our service.


## [GraphQL Readme](golang/internal/service/transport/graphql/README.md)


## TBD
- Merge checks(Gitlab pipeline for linters, unit and integration tests)
- Deployment
