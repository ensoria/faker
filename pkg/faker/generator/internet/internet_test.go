package internet_test

import (
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
	global := &provider.Global{
		Internets: global.CreateInternets(),
	}

	inet := internet.New(coreRand, global)

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

		It("Tld should return a tld", func() {
			r := inet.Tld()

			testutil.Output("Internet.Tld", r)
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

	It("TODO: Slug", func() {
		Skip("Slug: come back when Lorem is done")
	})

	It("TODO: Url", func() {
		Skip("Url: come back when Lorem is done")
	})

	Describe("Network", func() {
		It("Ipv4 should return a random ipv4 address", func() {
			r := inet.Ipv4()
			Expect(r).To(MatchRegexp(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`))
		})

		It("LocalIpv4 should return a random local ipv4 address", func() {
			r := inet.LocalIpv4()
			Expect(r).To(MatchRegexp(`(^10\.)|(^172\.1[6-9]\.)|(^172\.2[0-9]\.)|(^172\.3[0-1]\.)|(^192\.168\.)`))
		})

		It("Ipv6 should return a random ipv6 address", func() {
			r := inet.Ipv6()
			Expect(r).To(MatchRegexp(`^([0-9a-fA-F]{0,4}:){7}[0-9a-fA-F]{0,4}$`))
		})

		It("MacAddress should return a random mac address", func() {
			r := inet.MacAddress()
			Expect(r).To(MatchRegexp(`^([0-9a-fA-F]{2}:){5}[0-9a-fA-F]{2}$`))
		})
	})
})
