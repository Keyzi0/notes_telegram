package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Keyzi0/notes_telegram/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func main() {
	fmt.Println("Hello " + os.Getenv("BOT_TOKEN"))
	cfg := getConfig()
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func getConfig() models.Config {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	config.WithOptions(config.ParseEnv)

	// add driver for support yaml content
	config.AddDriver(yaml.Driver)

	if err = config.LoadFiles("config.yml"); err != nil {
		panic(err)
	}
	var configs models.Config
	config.BindStruct("bot", &configs)
	fmt.Println(fmt.Sprintf("config: %v", configs))
	return configs
}
