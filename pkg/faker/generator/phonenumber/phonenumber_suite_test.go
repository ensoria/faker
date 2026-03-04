package phonenumber_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPhoneNumber(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PhoneNumber Suite")
}
