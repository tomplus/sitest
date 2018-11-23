FROM golang:1.10-alpine

RUN apk add --no-cache git

WORKDIR /go/src/sitest
COPY *.go /go/src/sitest/

RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /go/bin/
COPY --from=0 /go/bin/sitest /go/bin/sitest

EXPOSE 8080

# nobody
USER 65534

CMD ["./sitest"]
