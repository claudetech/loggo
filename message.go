package loggo

import (
	"bytes"
	"github.com/mgutz/ansi"
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
	color      bool
	tpl        *template.Template
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

func (m *Message) String() string {
	buffer := bytes.NewBufferString("")
	err := m.tpl.Execute(buffer, m)
	if err != nil {
		return m.Content
	}
	str := buffer.String()
	if m.color {
		str = ansi.Color(str, Colors[m.Level])
	}
	return str
}

func (m *Message) TimeStr() string {
	return m.Time.Format(m.dateFormat)
}
