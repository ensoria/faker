package medical_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLorem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Medical Suite")
}
