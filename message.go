package loggo

import (
	"bytes"
	"fmt"
	"github.com/mgutz/ansi"
	"strings"
	"text/template"
	"time"
)

type Message struct {
	Name       string
	Level      Level
	Content    interface{}
	Time       time.Time
	File       string
	Line       int
	FuncName   string
	dateFormat string
	padding    bool
	color      bool
	tpl        *template.Template
}

func (m *Message) NameUp() string {
	return strings.ToUpper(m.Name)
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
		return fmt.Sprintf("%s\n", m.Content)
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
