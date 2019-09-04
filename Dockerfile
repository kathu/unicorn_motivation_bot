FROM golang:latest

ADD . /go/src/github.com/kathu/unicorn_motivation_bot

RUN go get -u "gopkg.in/telegram-bot-api.v4"

RUN go install github.com/kathu/unicorn_motivation_bot

ENTRYPOINT [ "/go/bin/unicorn_motivation_bot" ]

EXPOSE 80/tcp 3300