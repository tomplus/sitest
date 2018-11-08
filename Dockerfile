FROM golang:1.10

WORKDIR /go/src/sitest
COPY *.go /go/src/sitest/

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

# nobody
USER 65534

CMD ["sitest"]
