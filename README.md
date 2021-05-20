# go-msa

```
go mod tidy
```

```
go run main.go
```

Curl Command for testing
```
curl -G 'http://localhost:8080/user/1234'
```

To build the Docker Image
```
docker build -t example.com/go-msa:v1 .
```

To run image
```
docker run -p 8080:8080 -it example.com/go-msa:v1
```

kubectl deployment
```
kubectl create -f db-service.yaml,db-deployment.yaml,go-msa-service.yaml,claim0-volumeclaim.yaml,go-msa-deployment.yaml
```