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
		tpl, _ = template.New("foo").Parse(defaultFormat)
		msg = &Message{
			Name:       name,
			Level:      Debug,
			Content:    content,
			Time:       dummyTime(),
			dateFormat: defaultDateFormat,
			tpl:        tpl,
		}
	})

	It("should format message", func() {
		Expect(msg.String()).To(Equal(getString(tpl, msg)))
	})

	It("should respect template", func() {
		msg.tpl, _ = template.New("foo").Parse("{{.Content}}")
		Expect(msg.String()).To(Equal(content))
	})

	It("should use date format", func() {
		f := "Jan 2, 2006 at 15:04pm (MST)"
		msg.dateFormat = f
		Expect(msg.String()).To(ContainSubstring(msg.Time.Format(f)))
	})

	It("should work with padding", func() {
		msg.padding = true
		msg.tpl, _ = template.New("foo").Parse("{{.LevelStr}}:")
		Expect(msg.String()).To(Equal("DEBUG  :"))
	})
})
