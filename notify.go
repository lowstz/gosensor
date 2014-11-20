package main

import (
	"github.com/andybons/hipchat"
)

type notification interface {
	send()
}

type Hipchat struct {
	Token   string `json:"-"`
	Content string `json:content`
	From    string `json:from`
	To      string `json:to`
	Color   string `json:color`
}

type Email struct{}
type SMS struct{}

func (hc Hipchat) send() {
	c := hipchat.Client{AuthToken: hc.Token}
	req := hipchat.MessageRequest{
		RoomId:        hc.To,
		From:          hc.From,
		Message:       hc.Content,
		Color:         hc.Color,
		MessageFormat: "text",
		Notify:        true,
	}
	debug.Println(showCallerName(), req)

	if err := c.PostMessage(req); err != nil {
		errl.Println(showCallerName(), "Excepted no error, but got", err)
	} else {
		info.Println(showCallerName(), "hipchat send message successfully.")
	}
}

func (email Email) send() {

}

func (sms SMS) send() {

}

func SendNotification(n notification) {
	n.send()
}
