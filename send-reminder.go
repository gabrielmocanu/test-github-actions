package main

import (
    "fmt"
    "github.com/slack-go/slack"
)

func main() {
    api := slack.New("xoxb-2895394241809-2882746748658-3acoYBuUYpPwFIdUg4dWM5yN")

    channelID, timestamp, err := api.PostMessage(
        "C02SBBMHBL1",
        slack.MsgOptionText("Hello, it's time to write some golang", false),
    )

    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }

    fmt.Printf("Message send successfully to channel %s at %s", channelID, timestamp)
}
