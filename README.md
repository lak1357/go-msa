# Sample Microservice using go-kit and gorm

## Folder Structure 
```bash
├── Dockerfile
├── README.md
├── database
│   └── datbase.go
├── go.mod
├── go.sum
├── k8
│   ├── claim0-volumeclaim.yaml
│   ├── db-deployment.yaml
│   ├── db-service.yaml
│   ├── go-msa-deployment.yaml
│   └── go-msa-service.yaml
├── main.go
├── user
│   ├── endpoint.go
│   ├── logic.go
│   ├── repo.go
│   ├── reqresp.go
│   ├── server.go
│   ├── service.go
│   └── user.go
└── user.db
```

## Get packages
```bash
go mod tidy
```

## Run server
```bash
go run main.go
```

## Curl Command for testing

### Create User
```bash
curl --location --request POST 'http://localhost:8080/user' --header 'Content-Type: application/json' --data-raw '{"email" : "testing@gmail.com","password" : "xxxxxx"}'
```

### Get User
```bash
curl --location --request GET 'http://localhost:8080/user/{id}'
```

## To build the Docker Image
```bash
docker build -t example.com/go-msa:v1 .
```

## To run image
```bash
docker run -p 8080:8080 -it example.com/go-msa:v1
```

## kubectl deployment
```bash
kubectl create -f db-service.yaml,db-deployment.yaml,go-msa-service.yaml,claim0-volumeclaim.yaml,go-msa-deployment.yaml
```