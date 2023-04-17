package slack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type slackNotification struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji"`
}

// Notifier is the generic object to send slack notifications
type Notifier struct {
	SlackURL     string
	SlackChannel string
	Username     string
	IconEmoji    string
	Enable       bool
}

// InitializeNotifier creates a notifier instance
func InitializeNotifier(slackURL, slackChannel, username, emoji string, enable bool) *Notifier {
	return &Notifier{slackURL, slackChannel, username, emoji, enable}

}

// NotifySlack sends notification to slack channel created on the constructor
func (n *Notifier) NotifySlack(msg string) error {
	log.Print("Notification sent: ", msg)
	if n.Enable {
		url := n.SlackURL
		notify := slackNotification{
			Channel:   n.SlackChannel,
			Username:  n.Username,
			Text:      msg,
			IconEmoji: n.IconEmoji,
		}

		jsonValue, _ := json.Marshal(notify)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		_, err = client.Do(req)
		if err != nil {
			log.Print("Error on notification", err)
			return err
		}
	}
	return nil
}
