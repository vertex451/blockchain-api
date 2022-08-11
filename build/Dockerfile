FROM golang

ARG GIT_COMMIT
ENV GIT_COMMIT="$GIT_COMMIT"
LABEL commit="$GIT_COMMIT"

COPY ./build/artifacts/blockchain-api /
COPY ./.env /

EXPOSE 8080

ENTRYPOINT ["/blockchain-api"]
