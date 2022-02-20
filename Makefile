SHELL := /bin/bash

restart: stop start

start:
	@"${SHELL}" build/dev/scripts/registry.sh

	kind create cluster --config build/dev/kind_config.yml

	docker network connect "kind" "demo-registry" || true

	kubectl apply -f build/dev/registry_config_map.yml

	docker build --build-arg servicename=service -t localhost:5000/service:latest -f ./build/dev/Dockerfile .
	docker push localhost:5000/service:latest

	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
	kubectl wait --namespace ingress-nginx --for=condition=ready pod --selector=app.kubernetes.io/component=controller --timeout=90s

	kubectl apply -f build/dev/mysql.yml
	kubectl wait --namespace default --for=condition=ready pod --selector=app=mysql --timeout=90s

	kubectl apply -f build/dev/infra.yml

stop:
	kind delete cluster --name demo
	docker stop demo-registry || true
	docker rm demo-registry || true
