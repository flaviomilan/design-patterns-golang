package main

type message struct {
    messageType  string
    text string
}

func (m *message) setMessageType(messageType string) {
    m.messageType = messageType
}

func (m *message) getMessageType() string {
    return m.messageType
}

func (m *message) setText(text string) {
    m.text = text
}

func (m *message) getText() string {
    return m.text
}