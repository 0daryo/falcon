FROM golang:1.14.0-alpine3.11

RUN apk --update add --no-cache git
RUN go get -u github.com/oxequa/realize

CMD [ "realize", "start", "--run" ]