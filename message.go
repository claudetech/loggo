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
	padding    bool
}

func (m *Message) LevelStr() string {
	str := m.Level.String()
	if m.padding {
		for len(str) < 7 {
			str += " "
		}
	}
	return str
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
