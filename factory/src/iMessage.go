package main

type iMessage interface {
    setMessageType(name string)
    setText(text string)
    getMessageType() string
    getText() string
}