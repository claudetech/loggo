package loggo

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"text/template"
)

var _ = Describe("Message", func() {
	name := "FOO"
	content := "foo"

	var msg *Message
	var tpl *template.Template

	getString := func(t *template.Template, msg *Message) string {
		buffer := bytes.NewBufferString("")
		t.Execute(buffer, msg)
		return buffer.String()
	}

	BeforeEach(func() {
		msg = &Message{
			Name:       name,
			Level:      Debug,
			Content:    content,
			Time:       dummyTime(),
			dateFormat: defaultDateFormat,
		}
		tpl, _ = template.New("foo").Parse(defaultFormat)
	})

	It("should format message", func() {
		Expect(msg.Format(tpl)).To(Equal(getString(tpl, msg)))
	})

	It("should respect template", func() {
		tpl, _ := template.New("foo").Parse("{{.Content}}")
		Expect(msg.Format(tpl)).To(Equal(content))
	})

	It("should use date format", func() {
		f := "Jan 2, 2006 at 15:04pm (MST)"
		msg.dateFormat = f
		Expect(msg.Format(tpl)).To(ContainSubstring(msg.Time.Format(f)))
	})

	It("should work with padding", func() {
		msg.padding = true
		tpl, _ := template.New("foo").Parse("{{.LevelStr}}:")
		Expect(msg.Format(tpl)).To(Equal("DEBUG  :"))
	})
})
