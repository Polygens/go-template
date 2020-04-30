FIRSTRUN := $(shell [ -f ".git/hooks/commit-msg" ] && echo "file found")
K3D := $(shell command -v k3d 2> /dev/null)
K3D_CLUSTER := $(shell k3d l | grep local 2> /dev/null)
HOSTS := $(shell cat /etc/hosts | grep registry.local 2> /dev/null)
KAFKA := $(shell kubectl get svc | grep kafka 2> /dev/null)
export VERSION := $(shell git describe --tags --abbrev=0 &> /dev/null)-$(shell git rev-parse --abbrev-ref HEAD &> /dev/null)
export PROJECT := $(shell basename $(CURDIR))
export OUTPUT := $(shell echo $$HOME/go/bin/$(shell basename $(CURDIR)))
export DOCKER_BUILDKIT = 1

# Setup local dev environment
setup:
ifndef FIRSTRUN
	@echo "RUNNING INITIAL SETUP"

ifndef K3D
	@echo "k3d not found, installing..."
	@brew install k3d
endif
	
ifndef K3D_CLUSTER
	@echo "Creating a k3s cluster(local) and configuring kubectl"
	@-k3d create --enable-registry -n local --publish 8080:8080 --api-port 6550 --wait 120 && k3d get-kubeconfig -n local && \
	KUBECONFIG=~/.config/k3d/local/kubeconfig.yaml:~/.kube/config kubectl config view --flatten > ~/.kube/temp && \
	mv ~/.kube/temp ~/.kube/config
endif

ifndef KAFKA
	@echo "Adding bitname repo to helm repostiories"
	@helm repo add bitnami https://charts.bitnami.com/bitnami
	@echo "Installing kafka and zookeeper in the cluster"
	@helm install kafka bitnami/kafka --set externalAccess.enabled=true,externalAccess.service.type=LoadBalancer,externalAccess.service.port=19092,externalAccess.autoDiscovery.enabled=true,serviceAccount.create=true,rbac.create=true
	@echo "Configuring kaf"
	@kaf config add-cluster local -b local:19092
endif

	@echo "Setting up git hooks"
	@echo "#!/bin/bash\n\n. .github/commit.sh\nticket_prefix \$$1 \$$2" > .git/hooks/prepare-commit-msg
	@echo "#!/bin/bash\n\n. .github/commit.sh\nconventional_commit_validator \$$1" > .git/hooks/commit-msg
	@echo "#!/bin/bash\n\n. .github/commit.sh\ninsert_in_file README.md CONFIG \"\$$(yaml_table_config defaults.yaml)\"" > .git/hooks/pre-commit
	@echo "insert_in_file README.md KAFKA_INPUT \"\$$(kafka_input defaults.yaml)\"" >> .git/hooks/pre-commit
	@echo "insert_in_file README.md KAFKA_OUTPUT \"\$$(kafka_output defaults.yaml)\"" >> .git/hooks/pre-commit
	@chmod +x .git/hooks/prepare-commit-msg
	@chmod +x .git/hooks/commit-msg
	@chmod +x .git/hooks/pre-commit

ifndef HOSTS
	@echo "\nWARNING: Run the following lines to complete your setup:"
	@echo 'sudo -- sh -c "echo $$(kubectl get svc traefik -n kube-system -o jsonpath="{.status.loadBalancer.ingress[*].ip}") local >> /etc/hosts"'
	@echo 'sudo -- sh -c "echo 127.0.0.1 registry.local >> /etc/hosts"'; echo ""
endif
	@echo "COMPLETED INITIAL SETUP"
endif

# Local development
test:
	@go test ./...

build:
	@go build -ldflags="-X main.version=$$VERSION" -o $$OUTPUT

run: setup build
	@$$PROJECT 

# Containerized development
docker-test:
	@docker build --target=test -t $$PROJECT:test -q . && docker run $$PROJECT:test

docker-build:
	@docker build --target dev -t $$PROJECT -t registry.local:5000/$$PROJECT:latest --build-arg VERSION=$$VERSION --build-arg APP=$$PROJECT  .

setup-helm:
	@k3d start -n local &> /dev/null
	@kubectl config use-context local

helm: setup docker-build setup-helm
	@docker push registry.local:5000/$$PROJECT:latest
	@helm upgrade -i $$PROJECT ./charts --set image.repository=registry.local:5000/$$PROJECT --set image.pullPolicy=Always --wait --set ingress.enabled=true --recreate-pods

logs:
	@kubectl logs -l app.kubernetes.io/name=$$PROJECT -f

# Production
prod: setup setup-helm
	@docker build --target prod -t $$PROJECT -t registry.local:5000/$$PROJECT:latest --build-arg VERSION=$$VERSION --build-arg APP=$$PROJECT  .
	@docker push registry.local:5000/$$PROJECT:latest
	@helm upgrade -i $$PROJECT ./charts --set image.repository=registry.local:5000/$$PROJECT --set image.pullPolicy=Always --wait --set ingress.enabled=true --recreate-pods
