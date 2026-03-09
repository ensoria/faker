package internet_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/internet"
	"github.com/ensoria/faker/pkg/faker/provider"
	"github.com/ensoria/faker/pkg/faker/provider/global"
	"github.com/ensoria/faker/pkg/faker/testutil"
)

var _ = Describe("Internet", func() {
	coreRand := core.NewRand(util.RandSeed())
	g := &provider.Global{
		Internets: global.CreateInternets(),
		Lorems:    global.CreateLorems(),
	}

	inet := internet.New(coreRand, g)

	Describe("Email", func() {

		It("UserName should return a user name", func() {
			r := inet.UserName()

			Expect(r).To(MatchRegexp(`^[a-z][a-z0-9._-]{2,}$`))

			testutil.Output("Internet.UserName", r)
		})

		It("DomainWord should return a domain word", func() {
			r := inet.DomainWord()

			Expect(r).To(MatchRegexp(`^[a-z]+$`))
			testutil.Output("Internet.DomainWord", r)
		})

		It("TLD should return a tld", func() {
			r := inet.TLD()

			testutil.Output("Internet.TLD", r)
		})

		It("DomainName should return a domain name", func() {
			r := inet.DomainName()
			Expect(r).To(MatchRegexp(`^[a-z][a-z0-9._-]*\.[a-z]+$`))
			testutil.Output("Internet.DomainName", r)
		})

		It("Email should return an email", func() {
			r := inet.Email()
			Expect(r).To(MatchRegexp(`^[a-z][a-z0-9._-]*@[a-z][a-z0-9._-]*\.[a-z]+$`))
			testutil.Output("Internet.Email", r)
		})

		It("Password should return a random string between 8 to 20 length", func() {
			r := inet.Password()
			Expect(r).To(MatchRegexp(`^[\d\w]{8,20}$`))
		})
	})

	Describe("Slug", func() {
		It("should return hyphen-separated words", func() {
			r := inet.Slug(6, false)
			Expect(r).NotTo(BeEmpty())
			Expect(r).To(MatchRegexp(`^[a-z]+(-[a-z]+)*$`))
			testutil.Output("Internet.Slug", r)
		})

		It("should return the specified number of words when variableWordCount is false", func() {
			r := inet.Slug(3, false)
			parts := strings.Split(r, "-")
			Expect(parts).To(HaveLen(3))
		})

		It("should use default word count when nbWords is 0", func() {
			r := inet.Slug(0, false)
			parts := strings.Split(r, "-")
			Expect(parts).To(HaveLen(6))
		})

		It("should vary word count when variableWordCount is true", func() {
			// Run multiple times to verify it produces varying lengths
			counts := make(map[int]bool)
			for i := 0; i < 50; i++ {
				r := inet.Slug(6, true)
				parts := strings.Split(r, "-")
				counts[len(parts)] = true
			}
			// With variable word count, we should see more than one distinct length
			Expect(len(counts)).To(BeNumerically(">", 1))
		})
	})

	Describe("URL", func() {
		It("should return a valid URL", func() {
			r := inet.URL()
			Expect(r).To(MatchRegexp(`^https?://`))
			testutil.Output("Internet.URL", r)
		})

		It("should contain a domain name", func() {
			r := inet.URL()
			Expect(r).To(MatchRegexp(`https?://(www\.)?[a-z]+\.[a-z]+`))
		})
	})

	Describe("Network", func() {
		It("IPv4 should return a random ipv4 address", func() {
			r := inet.IPv4()
			Expect(r).To(MatchRegexp(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`))
		})

		It("LocalIPv4 should return a random local ipv4 address", func() {
			r := inet.LocalIPv4()
			Expect(r).To(MatchRegexp(`(^10\.)|(^172\.1[6-9]\.)|(^172\.2[0-9]\.)|(^172\.3[0-1]\.)|(^192\.168\.)`))
		})

		It("IPv6 should return a random ipv6 address", func() {
			r := inet.IPv6()
			Expect(r).To(MatchRegexp(`^([0-9a-fA-F]{0,4}:){7}[0-9a-fA-F]{0,4}$`))
		})

		It("MACAddress should return a random mac address", func() {
			r := inet.MACAddress()
			Expect(r).To(MatchRegexp(`^([0-9a-fA-F]{2}:){5}[0-9a-fA-F]{2}$`))
		})
	})
})
