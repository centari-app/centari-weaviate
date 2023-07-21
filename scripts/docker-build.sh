#docker buildx build --platform linux/amd64,linux/arm64 -t centari-weaviate:latest .

docker build --no-cache --platform linux/amd64 -t centari-weaviate:latest .