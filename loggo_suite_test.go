package loggo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLoggo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Loggo Suite")
}
