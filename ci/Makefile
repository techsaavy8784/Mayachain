
all: builder runner

builder:
	docker buildx build -t registry.gitlab.com/mayachain/mayanode:builder-v5 -f Dockerfile.builder .

runner:
	docker buildx build -t registry.gitlab.com/mayachain/mayanode:runner-base-v1 -f Dockerfile.runner .

push-builder:
	docker push registry.gitlab.com/mayachain/mayanode:builder-v5

push-runner:
	docker push registry.gitlab.com/mayachain/mayanode:runner-base-v1
