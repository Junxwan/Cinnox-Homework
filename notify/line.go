package notify

import (
	"Cinnox-Homework/cmd"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"net/http"
	"time"
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
		return fmt.Errorf("line webhook parse request error: %v", err)
	}

	for _, event := range events {
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			if _, err := b.Api.GetMessageQuota().Do(); err != nil {
				return fmt.Errorf("line message error: %v", err)
			}

			if _, err = b.Api.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
				apiErr := err.(*linebot.APIError)

				switch apiErr.Code {
				case http.StatusBadRequest:
					// 處理過期token，文件說是1分鐘但不保證且更改不通知
					if apiErr.Response.Message == "Invalid reply token" {
						return nil
					}

					// 由於上面用message當判斷怕有更改所以再加上時間做判斷
					if time.Now().Sub(event.Timestamp).Seconds() >= 60 {
						return nil
					}
				}

				return fmt.Errorf("line reply message error: %v, msg: %s", err, message.Text)
			}
		}
	}

	return nil
}
