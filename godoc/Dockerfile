FROM golang:1.20.2-alpine

RUN go install golang.org/x/tools/cmd/godoc@latest

EXPOSE 6060
CMD /go/bin/godoc -http=:6060
