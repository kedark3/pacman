# Pacman for assignment 1

# To run the project locally
```
git clone https://github.com/kedark3/pacman
podman(or docker) run -p 8080:8080 -v `pwd`:/pacman/static:Z --rm quay.io/kkulkarn/pacman
```
# What did I do in this Assignment? 

- I learned how to create a HTTP server to serve static content using GoLang.
- I learned how to build Go binaries and deploy it on OpenShift/Kubernetes. 
- In the first iteration I did not package static files and binary together hence I used initContainer to clone the repo and mount it in a volume to the pods.
- In 2nd and final iteration I learned how to generate static data with `vfsgen` that can be bundled into pacman's binary and serve it, see `data-generator` directory for the same.
- Another thing to notice here is the `Dockerfile` it uses a small container image `scratch` with multi-stage image reduction, and that helps make the final container image merely few MiBs vs 1+ GiB otherwise. 
- `k8s-configs` dir contains deployment, route and service object definitions you can use to deploy this in your cluster. The image is available on Quay.io.


# ArgoCD

Using ArgoCd CLI to create application as:
```sh
argocd app create -f application.yml

# OR 

kubectl create -f argocd/application.yml

```
