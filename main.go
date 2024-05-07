package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Emai)\n", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: SMS)\n", message)
}

type TelegramNotifier struct{}

func (TelegramNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: Telegram)\n", message)
}

type NotifierDecorator struct {
	core     *NotifierDecorator
	notifier Notifier
}

func (nd NotifierDecorator) Send(message string) {
	nd.notifier.Send(message)

	fmt.Println("nd.core: ", nd.core)
	if nd.core != nil {
		nd.core.Send(message)
	}
}

func (nd NotifierDecorator) Decorate(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{
		core:     &nd,
		notifier: notifier,
	}
}

func NewNotifierDecorator(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{notifier: notifier}
}

type NotificationService struct {
	notifier Notifier
}

func (notiService NotificationService) SendNotification(message string) {
	notiService.notifier.Send(message)
}

func main() {
	notifier := NewNotifierDecorator(EmailNotifier{}).
		Decorate(SMSNotifier{}).
		Decorate(TelegramNotifier{})

	notiService := NotificationService{notifier: notifier}

	notiService.SendNotification("Hello World")
}
