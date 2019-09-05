FROM golang:latest

ADD . /go/src/github.com/kathu/unicorn_motivation_bot

RUN go get -u "gopkg.in/telegram-bot-api.v4"

RUN go install github.com/kathu/unicorn_motivation_bot

RUN cp /go/src/github.com/kathu/unicorn_motivation_bot/token /go/src/github.com/kathu/unicorn_motivation_bot/proxy /go/bin

ENTRYPOINT [ "/go/bin/unicorn_motivation_bot" ]