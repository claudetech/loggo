package loggo

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"os"
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

		It("should write files", func() {
			path := "/tmp/foo.log"
			f, err := os.Create(path)
			Expect(err).To(BeNil())
			content, err := ioutil.ReadAll(f)
			Expect(err).To(BeNil())
			Expect(content).To(BeEmpty())
			f.Close()
			appender, err := NewFileAppender(path)
			Expect(err).To(BeNil())
			appender.Append(msg)
			Expect(appender.(io.Closer).Close()).To(BeNil())
			content, err = ioutil.ReadFile(path)
			Expect(err).To(BeNil())
			Expect(string(content)).To(Equal("foo"))
			Expect(os.Remove(path)).To(BeNil())
		})
	})
})
