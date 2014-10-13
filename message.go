package loggo

import (
	"bytes"
	"fmt"
	"github.com/mgutz/ansi"
	"strings"
	"text/template"
	"time"
)

// Message Represent a single log
type Message struct {
	// The name of the logger
	Name string
	// The level of the message
	Level Level
	// The content of the log
	Content interface{}
	// The time of the log
	Time time.Time
	// The file of the log call
	File string
	// The line number of the log call
	Line int
	// The function name of the log call
	FuncName   string
	dateFormat string
	padding    bool
	color      bool
	tpl        *template.Template
}

// Returns the logger name upper cased
func (m *Message) NameUp() string {
	return strings.ToUpper(m.Name)
}

// Formats the log level
func (m *Message) LevelStr() string {
	str := m.Level.String()
	if m.padding {
		for len(str) < 7 {
			str += " "
		}
	}
	return str
}

// Returns a formatted representation of the message
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

// Formats the time
func (m *Message) TimeStr() string {
	return m.Time.Format(m.dateFormat)
}
