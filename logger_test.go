package loggo

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"text/template"
	"time"
)

type dummyAppender struct {
	str string
}

func (d *dummyAppender) Append(msg *Message) {
	d.str += msg.String()
}

func dummyTime() time.Time {
	return time.Date(2009, time.November, 10, 15, 0, 0, 0, time.Local)
}

var _ = Describe("Logger", func() {
	var logger *Logger
	var appender *dummyAppender
	var tpl *template.Template

	name := "FOO"
	content := "foo"
	msg := &Message{
		Level:   Debug,
		Content: content,
	}

	getString := func(t *template.Template, msg *Message) string {
		msg.Name = name
		msg.Time = dummyTime()
		msg.dateFormat = defaultDateFormat
		buffer := bytes.NewBufferString("")
		t.Execute(buffer, msg)
		return buffer.String()
	}

	BeforeEach(func() {
		logger = New(name)
		appender = &dummyAppender{}
		logger.AddAppender(appender, 0)
		logger.SetNowFunc(dummyTime)
		logger.DisablePadding()
		tpl, _ = template.New("foo").Parse(defaultFormat)
	})

	It("should log", func() {
		logger.Debug(content)
		Expect(appender.str).To(Equal(getString(tpl, msg)))
	})

	It("should ignore when level is too low", func() {
		logger.Trace("foo")
		Expect(appender.str).To(BeEmpty())
	})

	It("should work with multiple appenders", func() {
		expected := getString(tpl, msg)
		expected = expected + expected
		logger.AddAppender(appender, 0)
		logger.Debug(content)
		Expect(appender.str).To(Equal(expected))
	})

	It("should ignore appenders if filter fails", func() {
		expected := getString(tpl, msg)
		filter := &MinLogLevelFilter{MinLevel: Warning}
		logger.AddAppenderWithFilter(appender, filter, 0)
		logger.Debug(content)
		Expect(appender.str).To(Equal(expected))
	})

	It("should respect template", func() {
		logger.SetFormat("{{.Content}}")
		logger.Debug(content)
		Expect(appender.str).To(Equal(content + "\n"))
	})

	It("should add padding", func() {
		logger.SetFormat("{{.LevelStr}}:")
		logger.EnablePadding()
		logger.Debug("foo")
		Expect(appender.str).To(Equal("DEBUG  :\n"))
	})

	It("should work with format", func() {
		logger.SetFormat("{{.Content}}")
		logger.Debugf("%s: %d + %.1f = %.1f", "Eq", 1, 1.1, 2.1)
		Expect(appender.str).To(Equal("Eq: 1 + 1.1 = 2.1\n"))

	})

	It("should work async", func() {
		logger.SetFormat("{{.Content}}")
		logger.AddAppender(appender, Async)
		logger.Debug("foo")
		Expect(appender.str).To(Equal("foo\n"))
		time.Sleep(10 * time.Millisecond) // not very safe way to check
		Expect(appender.str).To(Equal("foo\nfoo\n"))
	})

	It("should be destroyed", func() {
		logger.Destroy()
		Expect(logger.appenders).To(BeEmpty())
	})
})
