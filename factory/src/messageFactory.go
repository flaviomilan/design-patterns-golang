package main

import "fmt"

func getMessageType(messageType string) (iMessage, error) {
    if messageType == "sms" {
        return newSMS(), nil
    }
    if messageType == "email" {
        return newEmail(), nil
    }
    return nil, fmt.Errorf("Wrong message type passed")
}
