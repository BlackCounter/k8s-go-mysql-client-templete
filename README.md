## Usage

Deploy the temeplete application to kubernetes
```
docker build -t golang-mysql-client:local .
kubectl apply -f kubernetes/deployment.yaml
```

Get deployment's pod name
```
POD=$(kubectl get po | grep golang-mysql-client | awk '{print $1}')
```

Migration
```
kubectl exec $POD go run src/migration.go
```

Insert
```
kubectl exec $POD go run src/inser.go
```

Delete
```
kubectl exec $POD go run src/delete.go
```
