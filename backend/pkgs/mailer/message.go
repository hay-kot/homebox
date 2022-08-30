package mailer

import "net/mail"

type Message struct {
	Subject string
	To      mail.Address
	From    mail.Address
	Body    string
}

type MessageBuilder struct {
	subject string
	to      mail.Address
	from    mail.Address
	body    string
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{}
}

func (mb *MessageBuilder) Build() *Message {
	return &Message{
		Subject: mb.subject,
		To:      mb.to,
		From:    mb.from,
		Body:    mb.body,
	}
}

func (mb *MessageBuilder) SetSubject(subject string) *MessageBuilder {
	mb.subject = subject
	return mb
}

func (mb *MessageBuilder) SetTo(name, to string) *MessageBuilder {
	mb.to = mail.Address{
		Name:    name,
		Address: to,
	}
	return mb
}

func (mb *MessageBuilder) SetFrom(name, from string) *MessageBuilder {
	mb.from = mail.Address{
		Name:    name,
		Address: from,
	}
	return mb
}

func (mb *MessageBuilder) SetBody(body string) *MessageBuilder {
	mb.body = body
	return mb
}
