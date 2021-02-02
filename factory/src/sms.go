package main

type sms struct {
    message
}

func newSMS() iMessage {
    return &sms{
        message: message{
            messageType:  "SMS",
            text: "Message sent by SMS",
        },
    }
}