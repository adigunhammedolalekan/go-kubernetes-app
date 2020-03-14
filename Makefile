k8s:
	kubectl apply -f k8s/services.yml
	kubectl apply -f k8s/statefulsets.yml
	kubectl apply -f k8s/deployments.yml
	kubectl apply -f k8s/tls.yml
	kubectl apply -f k8s/ingress.yml

deploy:
	./build.sh
	k8s

make setup:
	# install ingress
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.30.0/deploy/static/mandatory.yaml
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.26.1/deploy/static/provider/cloud-generic.yaml

	# install cert-manager
	kubectl create namespace cert-manager
	kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v0.14.0/cert-manager.yaml