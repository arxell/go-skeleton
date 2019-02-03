FROM golang:1.11.3-alpine as builder
RUN apk update && apk add --no-cache git ca-certificates
COPY src/ $GOPATH/src/go-skeketon/src/
RUN set -ex \
    && ls $GOPATH/src/go-skeketon

WORKDIR $GOPATH/src/go-skeketon/
# RUN go get -u github.com/golang/dep/cmd/dep
# RUN dep ensure
RUN set -ex \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    && go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/app src/main.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]
