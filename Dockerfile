# STEP 1 build executable binary
FROM golang:alpine as builder

RUN apk update && apk add git && apk add ca-certificates
RUN adduser -D -g '' appuser
COPY . $GOPATH/src/github.com/goquotes
WORKDIR $GOPATH/src/github.com/goquotes
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o qoquotes .

# STEP 2 build a small image
# start from scratch
#FROM scratch
FROM golang:alpine
# Copy our static executable
COPY --from=builder /go/src/github.com/goquotes/qoquotes /go/bin/goquotes
COPY --from=builder /go/src/github.com/goquotes/static /go/bin/static
COPY --from=builder /go/src/github.com/goquotes/templates /go/bin/templates
#COPY --from=builder /go/bin/hello /go/bin/goquotes
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#EXPOSE 3000
USER appuser
#RUN chmod 700 /go/bin/quotes/target
ENTRYPOINT ["/go/bin/goquotes"]