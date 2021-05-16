# go-msa

```
go mod tidy
```

```
go run main.go
```


To build the Docker Image
```
docker build -t example.com/go-msa:v1 .
```

To run image
```
docker run -p 8080:8080 -it example.com/go-msa:v1
```