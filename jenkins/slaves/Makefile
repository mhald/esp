PROJECT = endpoints-jenkins
VERSION = 0.8
TOOLS_BUCKET = endpoints-tools

# Note: The build directory is the root of the istio/test-infra repository, not ./
image:
	docker build -t debian-9 --build-arg TOOLS_BUCKET="${TOOLS_BUCKET}" -f debian-9.Dockerfile ../../
	docker tag debian-9 gcr.io/$(PROJECT)/debian-9:$(VERSION)
	docker tag debian-9 gcr.io/$(PROJECT)/debian-9:latest


push: image
	docker push gcr.io/$(PROJECT)/debian-9:$(VERSION)
	docker push gcr.io/$(PROJECT)/debian-9:latest


.PHONY: image push
