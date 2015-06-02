package loggo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Level", func() {
	Describe("LevelFromString", func() {
		It("should return the correct value", func() {
			cases := map[string]Level{
				"trace":   Trace,
				"debug":   Debug,
				"info":    Info,
				"warning": Warning,
				"error":   Error,
				"fatal":   Fatal,
			}
			for in, out := range cases {
				Expect(LevelFromString(in)).To(Equal(out))
			}
		})
	})
})
