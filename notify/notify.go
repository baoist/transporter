package notify

import (
	"github.com/0xAX/notificator"
)

var notification *notificator.Notificator

func Notify(title, message string) {
	notify(title, message)
}

func notify(title, message string) {
	notification = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "Transporter",
	})

	notification.Push(title, message, "", notificator.UR_CRITICAL)
}
