package loggo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Filter", func() {

	Describe("LogLevelFilter", func() {
		var filter *LogLevelFilter
		var msg *Message

		BeforeEach(func() {
			filter = &LogLevelFilter{MinLevel: Info}
			msg = &Message{
				Name:       "foo",
				Level:      Debug,
				Content:    "bar",
				Time:       time.Now(),
				dateFormat: defaultDateFormat,
			}
		})

		It("should return true when log level is high enough", func() {
			Expect(filter.ShouldLog(msg)).To(BeFalse())
			msg.Level = Info
			Expect(filter.ShouldLog(msg)).To(BeTrue())
			msg.Level = Warning
			Expect(filter.ShouldLog(msg)).To(BeTrue())
		})
	})
})
