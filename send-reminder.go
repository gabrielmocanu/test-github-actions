package main

import (
    "fmt"
    "os"
    "github.com/slack-go/slack"
)

func main() {
    api := slack.New(os.Getenv("OAUTH_BOT_TOKEN"))

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
