package pd_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gock "gopkg.in/h2non/gock.v1"
)

var _ = BeforeSuite(func() {
	defer gock.OffAll()
	defer gock.DisableNetworking()
})

func TestPD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PD Backend")
}
