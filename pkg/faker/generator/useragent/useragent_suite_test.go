package useragent_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUserAgent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UserAgent Suite")
}
