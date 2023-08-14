package teams

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type teamsBody struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type teamsContent struct {
	Schema  string      `json:"$schema"`
	Type    string      `json:"type"`
	Version string      `json:"version"`
	Body    []teamsBody `json:"body"`
}

type teamsAttachment struct {
	ContentType string       `json:"contentType"`
	ContentURL  string       `json:"contentURL"`
	Content     teamsContent `json:"content"`
}

type teamsNotification struct {
	Type        string            `json:"type"`
	Attachments []teamsAttachment `json:"attachments"`
	Username    string            `json:"username"`
	Text        string            `json:"text"`
	IconEmoji   string            `json:"icon_emoji"`
}

// Notifier is the generic object to send teams notifications
type TeamsNotifier struct {
	TeamsWebHookURL string
	Username        string
	IconEmoji       string
	Enable          bool
}

// InitializeNotifier creates a notifier instance
func InitializeNotifier(TeamsWebHookURL, username, emoji string, enable bool) *TeamsNotifier {
	return &TeamsNotifier{TeamsWebHookURL, username, emoji, enable}
}

// NotifyTeams sends notification to teams channel created on the constructor
func (n *TeamsNotifier) NotifyTeams(msg string) error {
	log.Print("Notification sent: ", msg)
	if n.Enable {
		url := n.TeamsWebHookURL
		notify := teamsNotification{
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
