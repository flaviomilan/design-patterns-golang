package main

type email struct {
    message
}

func newEmail() iMessage {
    return &email{
        message: message{
            message_type:  "E-mail",
            text: "Message sent by E-mail",
        },
    }
}