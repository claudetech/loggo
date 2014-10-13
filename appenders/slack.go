package appenders

import (
	"bytes"
	"encoding/json"
	"github.com/claudetech/loggo"
	"net/http"
)

type SlackAppender struct {
	url     string
	client  *http.Client
	message *slackMessage
}

type slackMessage struct {
	Icon     string `json:"icon_emoji,omitempty"`
	Channel  string `json:"channel,omitempty"`
	Username string `json:"username,omitempty"`
	Text     string `json:"text"`
}

func NewSlackAppender(url string, username string, icon string, channel string) *SlackAppender {
	client := http.DefaultClient
	slackMessage := &slackMessage{
		Username: username,
		Icon:     icon,
		Channel:  channel,
	}
	return &SlackAppender{
		url:     url,
		client:  client,
		message: slackMessage,
	}
}
func (s *SlackAppender) Append(msg *loggo.Message) {
	s.message.Text = msg.String()
	body, err := json.Marshal(s.message)
	if err != nil {
		return
	}
	s.client.Post(s.url, "application/json", bytes.NewReader(body))
}
