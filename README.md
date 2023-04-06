# NGINX Operator

![](https://img.shields.io/badge/platform-kubernetes-blue)
![](https://img.shields.io/badge/operator-shell%20script-brightgreen)
![](https://img.shields.io/github/v/release/amirhnajafiz/nginx-operator)

Nginx operator for Kubernetes. This operator watches the changes of ```ConfigMaps``` and
rollouts a new deployment whenever you release a new ```ConfigMap``` for you ```nginx``` application.

## Start

In your cluster run the script:

```shell
chmod +x ./operator.sh && ./operator.sh
```

### Sample

A sample of ```nginx``` configmap:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: sample
data:
  index.html: hello world
```

### Test

Now if you apply this ```ConfigMap```, operator will automatically releases a new deployment:

```shell
kubectl apply -f nginx.yml
```
