sudo curl -sfL https://get.k3s.io |INSTALL_K3S_SKIP_DOWNLOAD=true K3S_KUBECONFIG_MODE=644 sh -
kubectl delete -f /var/lib/rancher/k3s/server/manifests/traefik.yaml
kubectl apply -f ./kong-ingress.yaml
docker run -p 1337:1337 -e "TOKEN_SECRET=723rebwher982h3rh239rbb8hvysuh" -e "NODE_ENV=production" --name konga pantsel/konga
sudo docker run --privileged -d -v /home/double/App/docker/data/rancher:/var/lib/rancher -p 8888:80 -p 8443:443 rancher/rancher
