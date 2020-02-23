FROM golang:1.13 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/tomles ./cmd/tomles/

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/tomles /

ENTRYPOINT ["/tomles"]
