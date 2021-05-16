FROM golang:latest

RUN mkdir -p /go/src/example.com/go-msa

WORKDIR /go/src/example.com/go-msa

COPY . /go/src/example.com/go-msa

RUN go install example.com/go-msa

CMD /go/bin/example.com/go-msa

EXPOSE 8080