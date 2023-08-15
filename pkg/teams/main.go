package teams

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type ContentBody struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type AttachmentContent struct {
	Schema  string        `json:"$schema"`
	Type    string        `json:"type"`
	Version string        `json:"version"`
	Body    []ContentBody `json:"body"`
}

type Attachment struct {
	ContentType string            `json:"contentType"`
	ContentURL  *string           `json:"contentURL"`
	Content     AttachmentContent `json:"content"`
}

type Notification struct {
	Type        string       `json:"type"`
	Attachments []Attachment `json:"attachments"`
}

// Notifier is the generic object to send teams notifications
type TeamsNotifier struct {
	TeamsWebHookURL string
	Enable          bool
}

// InitializeNotifier creates a notifier instance
func InitializeNotifier(TeamsWebHookURL string, Enable bool) *TeamsNotifier {
	return &TeamsNotifier{TeamsWebHookURL, Enable}
}

func MarshalledJSONNotification(message string) ([]byte, error) {
	contentBody := ContentBody{
		Type: "TextBlock",
		Text: message,
	}

	attachmentContent := AttachmentContent{
		Schema:  "http://adaptivecards.io/schemas/adaptive-card.json",
		Type:    "AdaptiveCard",
		Version: "1.2",
		Body:    append(make([]ContentBody, 1), contentBody),
	}

	attachment := Attachment{
		ContentType: "application/vnd.microsoft.card.adaptive",
		Content:     attachmentContent,
	}

	notification := Notification{
		Type:        "message",
		Attachments: append(make([]Attachment, 1), attachment),
	}

	return json.Marshal(notification)
}

// NotifyTeams sends notification to teams channel created on the constructor
func (n *TeamsNotifier) NotifyTeams(message string) error {
	if n.Enable {
		jsonValue, _ := MarshalledJSONNotification(message)

		req, err := http.NewRequest("POST", n.TeamsWebHookURL, bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		_, err = client.Do(req)
		if err != nil {
			log.Print("Error on notification", err)
			return err
		}
		log.Print("Notification sent: ", message)
	} else {
		log.Print("Skipping Microsoft Teams notification: ", message)
	}
	return nil
}
