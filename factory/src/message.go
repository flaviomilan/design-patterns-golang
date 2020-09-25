package main

type message struct {
    message_type  string
    text string
}

func (m *message) setMessageType(message_type string) {
    m.message_type = message_type
}

func (m *message) getMessageType() string {
    return m.message_type
}

func (m *message) setText(text string) {
    m.text = text
}

func (m *message) getText() string {
    return m.text
}