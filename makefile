export FIRSTRUN := $(shell [ -f ".git/hooks/commit-msg" ] && echo "true" || echo "false")
export VERSION := $(shell git describe --tags --abbrev=0 &> /dev/null):$(shell git rev-parse --abbrev-ref HEAD &> /dev/null)
export PROJECT := $(shell basename $(CURDIR))
export OUTPUT := $(shell echo $$HOME/go/bin/$(shell basename $(CURDIR)))
export DOCKER_BUILDKIT = 1

# Setup local dev environment
setup:
ifeq ($(FIRSTRUN), false)
	# install k3d/k3s
	brew install k3d
	# Create a k3s cluster, configure kubectl
	-k3d create --enable-registry -n k3s --publish 8080:8080 --api-port 6550 --wait 120 && k3d get-kubeconfig -n k3s && \
	KUBECONFIG=~/.config/k3d/k3s/kubeconfig.yaml:~/.kube/config kubectl config view --flatten > ~/.kube/temp && \
	mv ~/.kube/temp ~/.kube/config
	# Confgure jira ticket and conventional commit hooks
	echo "#!/bin/bash\n\n. .github/commit.sh\nticket_prefix \$$1 \$$2" > .git/hooks/prepare-commit-msg
	echo "#!/bin/bash\n\n. .github/commit.sh\nconventional_commit_validator \$$1" > .git/hooks/commit-msg
	# Setup hosts
	@echo "\n\nWARNING: Run the following lines to complete your setup:"
	@echo 'sudo -- sh -c "echo $$(kubectl get svc traefik -n kube-system -o jsonpath="{.status.loadBalancer.ingress[*].ip}") k3s >> /etc/hosts"'
	@echo 'sudo -- sh -c "echo 127.0.0.1 registry.local >> /etc/hosts"'; echo "\n\n"endif
endif

# Testing
test:
	@docker build --target=test -t $$PROJECT:test -q . && docker run $$PROJECT:test

# Development
build:
	go build -ldflags="-X main.version=$$VERSION" -o $$OUTPUT

run: setup build
	@$$PROJECT 

dev-build:
	@docker build --target dev -t $$PROJECT -t registry.local:5000/$$PROJECT:latest --build-arg VERSION=$$VERSION --build-arg APP=$$PROJECT  .

setup-helm:
	k3d start -n k3s
	kubectl config use-context k3s

helm: setup dev-build setup-helm
	@docker push registry.local:5000/$$PROJECT:latest
	helm upgrade -i $$PROJECT ./charts --set image.repository=registry.local:5000/$$PROJECT --set image.pullPolicy=Always --wait --set ingress.enabled=true

logs:
	@kubectl logs -l app.kubernetes.io/name=$$PROJECT -f
