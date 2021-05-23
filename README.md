# go-msa

```
go mod tidy
```

```
go run main.go
```

Curl Command for testing

Create User
```
curl --location --request POST 'http://localhost:8080/user' --header 'Content-Type: application/json' --data-raw '{"email" : "testing@gmail.com","password" : "xxxxxx"}'
```

Get User
```
curl --location --request GET 'http://localhost:8080/user/{id}'
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