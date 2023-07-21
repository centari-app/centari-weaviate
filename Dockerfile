# Dockerfile for development purposes.
# Read docs/development.md for more information
# vi: ft=dockerfile


###############################################################################
# Base build image
FROM golang:1.20-alpine AS build_base
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/weaviate/weaviate
ENV GO111MODULE=on
# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

###############################################################################
# This image builds the weaviate server
FROM build_base AS server_builder
ARG TARGETARCH
ARG GITHASH="unknown"
ARG EXTRA_BUILD_ARGS=""
COPY . .
RUN GOOS=linux GOARCH=$TARGETARCH go build $EXTRA_BUILD_ARGS \
      -ldflags '-w -extldflags "-static" -X github.com/weaviate/weaviate/usecases/config.GitHash='"$GITHASH"'' \
      -o /weaviate-server ./cmd/weaviate-server

###############################################################################
# This creates an image that can be used to fake an api for telemetry acceptance test purposes
FROM build_base AS telemetry_mock_api
COPY . .
ENTRYPOINT ["./tools/dev/telemetry_mock_api.sh"]

###############################################################################
# Weaviate (no differentiation between dev/test/prod - 12 factor!)
FROM alpine AS weaviate
ENTRYPOINT ["/bin/weaviate"]
COPY --from=server_builder /weaviate-server /bin/weaviate
RUN apk add --no-cache --upgrade ca-certificates openssl
RUN mkdir ./modules
ENV AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED=true
ENV CLUSTER_HOSTNAME="node1"
ENV DEFAULT_VECTORIZER_MODULE="text2vec-openai"
ENV ENABLE_MODULES=text2vec-openai,qna-openai
ENV PERSISTENCE_DATA_PATH="/var/data"
CMD [ "--host", "0.0.0.0", "--port", "8080", "--scheme", "http"]
EXPOSE 8080