package loggo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Filter", func() {
	var msg *Message
	var filter Filter

	BeforeEach(func() {
		msg = &Message{
			Name:       "foo",
			Level:      Debug,
			Content:    "bar",
			Time:       time.Now(),
			dateFormat: defaultDateFormat,
		}
	})

	Describe("MinLogLevelFilter", func() {
		BeforeEach(func() {
			filter = &MinLogLevelFilter{MinLevel: Info}

		})

		It("should return true when log level is high enough", func() {
			Expect(filter.ShouldLog(msg)).To(BeFalse())
			msg.Level = Info
			Expect(filter.ShouldLog(msg)).To(BeTrue())
			msg.Level = Warning
			Expect(filter.ShouldLog(msg)).To(BeTrue())
		})
	})

	Describe("MaxLogLevelFilter", func() {
		BeforeEach(func() {
			filter = &MaxLogLevelFilter{MaxLevel: Info}
		})

		It("should return true when log level is high enough", func() {
			Expect(filter.ShouldLog(msg)).To(BeTrue())
			msg.Level = Info
			Expect(filter.ShouldLog(msg)).To(BeTrue())
			msg.Level = Warning
			Expect(filter.ShouldLog(msg)).To(BeFalse())
		})
	})
})
