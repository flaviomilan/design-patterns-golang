package main

import "fmt"

func main() {
    sms, _ := getMessageType("sms")
    email, _ := getMessageType("email")

    printDetails(sms)
    printDetails(email)
}

func printDetails(m iMessage) {
    fmt.Printf("Method: %s", m.getMessageType())
    fmt.Println()
    fmt.Printf("Text: %s", m.getText())
    fmt.Println()
}
