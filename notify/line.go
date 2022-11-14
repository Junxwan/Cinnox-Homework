package notify

import (
	"Cinnox-Homework/cmd"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"net/http"
)

type Bot struct {
	Api *linebot.Client
}

func New(conf cmd.Line) (*Bot, error) {
	line, err := linebot.New(conf.Secret, conf.Token)
	if err != nil {
		return nil, fmt.Errorf("new line bot error: %v", err)
	}

	bot := new(Bot)
	bot.Api = line

	return bot, nil
}

// 處理Line message Webhook
func (b *Bot) ConsumeMessage(req *http.Request) error {
	events, err := b.Api.ParseRequest(req)

	if err != nil {
		return err
	}

	for _, event := range events {
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			if _, err := b.Api.GetMessageQuota().Do(); err != nil {
				return err
			}

			if _, err = b.Api.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
				return err
			}
		}
	}

	return nil
}
