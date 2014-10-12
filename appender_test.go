package loggo

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"text/template"
)

var _ = Describe("Appender", func() {
	Describe("writerAppender", func() {

		content := "foo"
		tpl, _ := template.New("foo").Parse("{{.Content}}")
		msg := &Message{
			Level:   Debug,
			Content: content,
			tpl:     tpl,
		}

		It("should write to Writer", func() {
			w := bytes.NewBufferString("")
			appender := NewWriterAppender(w)
			appender.Append(msg)
			Expect(w.String()).To(Equal("foo"))
		})
	})
})
