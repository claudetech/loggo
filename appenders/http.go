package loggo_appenders

import (
	"bytes"
	"encoding/json"
	"github.com/claudetech/loggo"
	"net/http"
)

type HttpAppender struct {
	requestGenerator func(loggo.Message) (*http.Request, error)
	url              string
}

func defaultRequestGenerator(url string) func(loggo.Message) (*http.Request, error) {
	return func(msg loggo.Message) (r *http.Request, err error) {
		var body []byte
		if body, err = json.Marshal(map[string]string{"text": msg.String()}); err != nil {
			return
		}
		if r, err = http.NewRequest("POST", url, bytes.NewReader(body)); err != nil {
			return
		}
		r.Header.Add("Content-Type", "application/json")
		return
	}
}

// func (h *HttpAppender) Append(msg *Message) {
// 	r, err := defaultRequestGenerator(h.url)
// 	if err != nil {
// 		return
// 	}
// 	_, _ = http.Client.Do(r)
// }
