FROM golang:1.15

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run",  "."]