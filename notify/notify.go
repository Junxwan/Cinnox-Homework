package notify

import (
	"net/http"
	"time"
)

type INotify interface {
	Webhook(req *http.Request) error
	Send(userId, message string) error
	Sends(userIds []string, message string) error
	Broadcast(message string) error
}

type IModel interface {
	CreateUser(id, name string, created time.Time) error
	CreateMessage(userId, message string, created time.Time) error
}
