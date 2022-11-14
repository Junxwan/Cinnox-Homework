package notify

import (
	"net/http"
	"time"
)

type INotify interface {
	Webhook(req *http.Request) error
}

type IModel interface {
	CreateUser(id, name string, created time.Time) error
	CreateMessage(userId, message string, created time.Time) error
}
