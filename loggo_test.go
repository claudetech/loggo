package loggo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("loggo", func() {
	Describe("loggers", func() {
		It("should save created loggers", func() {
			_ = New("foo")
			Expect(Get("foo")).ToNot(BeNil())
		})

		It("should remove loggers on destroy", func() {
			_ = Get("foo").Destroy()
			Expect(Get("foo")).To(BeNil())
		})
	})
})
