package notify

import "net/http"

type INotify interface {
	Webhook(req *http.Request) error
}