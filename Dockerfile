FROM golang:1.17-alpine3.13
WORKDIR /go/src/github.com/Keyzi0/notes_telegram
RUN apk add --no-cache make git
COPY . .
RUN go build -o /main
CMD ["/main"]