package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strings"

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
		if regexp.MustCompile("спать|усну|засыпаю|сонли").MatchString(strings.ToLower(update.Message.Text)) {
			resps := [...]string{
				"Сон для слабых",
				"Пока ты спишь - враг качается!",
				"Eat, sleep, hack, repeat"}

			response = resps[rand.Intn(len(resps)-1)]
		} else if regexp.MustCompile("устал|сил*нет|нет*сил|заебал.?с").MatchString(strings.ToLower(update.Message.Text)) {
			resps := [...]string{
				"Знаешь, ты можешь добиться большего... Не сдавайся",
				"Даже если ты не веришь в единорогов.. Единороги верят в тебя. Не подведи их!",
				"Весна юности все еще не закончена! Рано терять надежду!",
				"Самое важное для ниндзя — не количество освоенных техник. Самое важное для него это — никогда не сдаваться!"}

			response = resps[rand.Intn(len(resps)-1)]
		} else if regexp.MustCompile("дибил|дурак|тупой|идиот|зря\\s|долбое|тупиц|дурен").MatchString(strings.ToLower(update.Message.Text)) {
			resps := [...]string{
				"Человеку свойственно ошибаться, и он пользуется этим свойством часто и с удовольствием",
				"Те, кто не понимают себя, проигрывают"}

			response = resps[rand.Intn(len(resps)-1)]
		} else if regexp.MustCompile("жаль|жалко|увы").MatchString(strings.ToLower(update.Message.Text)) {
			resps := [...]string{
				"Уже слишком поздно о чем-то сожалеть. Реальность безжалостна в своем движении вперед",
				"Весна юности все еще не закончена! Рано терять надежду!"}

			response = resps[rand.Intn(len(resps)-1)]
		} else {
			command := update.Message.Command()
			if command != "" {
				fmt.Printf("command: %s", command)
				if regexp.MustCompile("motivation").MatchString(strings.ToLower(command)) {
					resps := [...]string{
						"За моей спиной крылья, имя которым — настойчивость. Порой настойчивость может принять форму крыльев и сделать даже невозможное возможным!",
						"Все люди делятся на гениев и тех, кто всего добивается своим трудом",
						"Нет ничего невозможного, если вкладывать в дело душу",
						"Если не пытаться что-то изменить, то ничего и не изменится",
						"Люди никогда не протянут руку помощи тому, кто ничего не делает и бежит от трудностей. Пока ты не сдаёшься, всегда есть надежда на спасение",
						"Направляй свои мысли куда следует, иначе они направят тебя, куда не следует",
						"Люди меняются или же умирают, так и не изменившись. Одно из двух",
						"Знаешь, ты можешь добиться большего... Не сдавайся",
						"Даже если ты не веришь в единорогов.. Единороги верят в тебя. Не подведи их!",
						"Усердие всегда приносит свои плоды!",
						"Врождённый талант может не принести счастья. Счастливы те, кто верят в себя до самого конца и упорно трудятся",
						"Если есть упорство, тебе всё по плечу",
						"Весна юности все еще не закончена! Рано терять надежду!",
						"Если ты что-то начал, то должен обязательно закончить, и нельзя останавливаться, пока не будет так, как надо, а если не получается, нужно начинать сначала, и так пока не добьёшься успеха",
						"Не жалуйся, если не собираешься сдаваться",
						"Самое важное для ниндзя — не количество освоенных техник. Самое важное для него это — никогда не сдаваться!",
						"Советую вам не следовать моим советам"}

					response = resps[rand.Intn(len(resps)-1)]
				}
			}

		}

		if len(response) != 0 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
