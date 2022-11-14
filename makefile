name = alexa
registry = registry.mbenaissa.dev

docker-build:
	docker buildx build --platform linux/arm/v7,linux/arm64 -t ${registry}/${name}:latest --push -f api/Dockerfile .

deploy-alexa:
	cd alexa && ask deploy