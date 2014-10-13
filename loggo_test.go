package loggo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"testing"
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

func BenchmarkStdoutLogger(b *testing.B) {
	stdout = ioutil.Discard
	b.StopTimer()
	logger := New("foo")
	logger.AddAppender(NewStdoutAppender(), EmptyFlag)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug("foo")
	}
	b.StopTimer()
}

func BenchmarkStdoutNotLogged(b *testing.B) {
	b.StopTimer()
	logger := New("foo")
	logger.SetLevel(Info)
	logger.AddAppender(NewStdoutAppender(), EmptyFlag)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug("foo")
	}
	b.StopTimer()
}

func BenchmarkFileLogger(b *testing.B) {
	b.StopTimer()
	logger := New("bar")
	filepath := "/tmp/loggo.log"
	appender, err := NewFileAppender(filepath)
	if err != nil {
		b.Fail()
	}
	logger.AddAppender(appender, EmptyFlag)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Debug("foo")
	}
	b.StopTimer()
	os.Remove(filepath)
}
