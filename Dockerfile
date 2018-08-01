
FROM golang:1.8

WORKDIR /go/src/app

COPY ./src/app .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
