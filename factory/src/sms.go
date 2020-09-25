package main

type sms struct {
    message
}

func newSMS() iMessage {
    return &sms{
        message: message{
            message_type:  "SMS",
            text: "Message sent by SMS",
        },
    }
}