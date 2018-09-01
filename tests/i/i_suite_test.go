package i_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "I")
}
