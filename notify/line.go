package notify

import (
	"Cinnox-Homework/cmd"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"net/http"
)

type Bot struct {
	Api *linebot.Client

	model IModel
}

func New(conf cmd.Line, model IModel) (*Bot, error) {
	line, err := linebot.New(conf.Secret, conf.Token)
	if err != nil {
		return nil, fmt.Errorf("new line bot error: %v", err)
	}

	bot := new(Bot)
	bot.Api = line
	bot.model = model

	return bot, nil
}

// 處理Line message Webhook
func (b *Bot) Webhook(req *http.Request) error {
	events, err := b.Api.ParseRequest(req)

	if err != nil {
		return fmt.Errorf("line webhook parse request error: %v", err)
	}

	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeFollow:
			return b.follow(event)
		case linebot.EventTypeMessage:
			return b.message(event)
		}
	}

	return nil
}

// 處理文字
func (b *Bot) message(event *linebot.Event) error {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		//msg := fmt.Sprintf("msg: %s\ntime: %s\nuser: %s", message.Text, event.Timestamp.String(), event.Source.UserID)
		//if _, err := b.Api.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
		//	apiErr := err.(*linebot.APIError)
		//
		//	switch apiErr.Code {
		//	case http.StatusBadRequest:
		//		// 處理過期token，文件說是1分鐘但不保證且更改不通知
		//		if apiErr.Response.Message == "Invalid reply token" {
		//			return nil
		//		}
		//
		//		// 由於上面用message當判斷怕有更改所以再加上時間做判斷
		//		if time.Now().Sub(event.Timestamp).Seconds() >= 60 {
		//			return nil
		//		}
		//	}
		//
		//	return fmt.Errorf("line reply message error: %v, msg: %s", err, message.Text)
		//}

		return b.model.CreateMessage(event.Source.UserID, message.Text, event.Timestamp)
	}

	return nil
}

// 新關注或解封鎖
func (b *Bot) follow(event *linebot.Event) error {
	user, err := b.Api.GetProfile(event.Source.UserID).Do()
	if err != nil {
		return err
	}

	return b.model.CreateUser(user.UserID, user.DisplayName, event.Timestamp)
}
