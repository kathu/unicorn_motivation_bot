FROM golang:latest
RUN mkdir /bot
ADD . /bot/
WORKDIR /bot
RUN go get -u "gopkg.in/telegram-bot-api.v4"

RUN go build -o main .

ENTRYPOINT [ "./main" ]