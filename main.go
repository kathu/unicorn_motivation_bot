package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

func main() {
	token, _ := ioutil.ReadFile("token")
	proxyStr, _ := ioutil.ReadFile("proxy")
	proxyUrl, err := url.Parse(string(proxyStr))

	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	bot, err := tgbotapi.NewBotAPIWithClient(string(token), client)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		var response string
		isSleep := regexp.MustCompile("спать|усну|засыпаю|сонли").MatchString(update.Message.Text)
		if isSleep {
			response = "Сон для слабых"
		}

		if len(response) != 0 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
