# ingress-nginx example

For convenience, some manifest files from [kubernetes/ingress-nginx](https://github.com/kubernetes/ingress-nginx/) repo are saved here:

 - Modified: [mandatory.yaml](https://github.com/kubernetes/ingress-nginx/blob/master/deploy/static/mandatory.yaml)

 - Unmodified: [cloud-generic.yaml](https://github.com/kubernetes/ingress-nginx/blob/master/deploy/static/provider/cloud-generic.yaml)



Install:

```
kubectl apply -f mandatory.yaml

kubectl apply -f cloud-generic.yaml

```


See if successful:

```
kubectl get all -n ingress-nginx
```


For more infomation, see [NGINX Ingress Controller Installation Guide](https://kubernetes.github.io/ingress-nginx/deploy/) maintained by K8s official site.