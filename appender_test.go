package loggo

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Appender", func() {
	Describe("writerAppender", func() {

		It("should write to Writer", func() {
			w := bytes.NewBufferString("")
			appender := NewWriterAppender(w)
			appender.Append("foobar", Debug)
			Expect(w.String()).To(Equal("foobar"))
		})
	})
})
