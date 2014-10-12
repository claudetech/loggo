package loggo

import (
	"bytes"
	"text/template"
	"time"
)

type Message struct {
	Name       string
	Level      Level
	Content    string
	Time       time.Time
	dateFormat string
}

func (m *Message) LevelStr() string {
	return m.Level.String()
}

func newMessage(level Level, content string) *Message {
	return &Message{
		Level:   level,
		Content: content,
	}
}

func (m *Message) Format(tpl *template.Template) string {
	buffer := bytes.NewBufferString("")
	err := tpl.Execute(buffer, m)
	if err != nil {
		return m.Content
	}
	return buffer.String()
}

func (m *Message) TimeStr() string {
	return m.Time.Format(m.dateFormat)
}
