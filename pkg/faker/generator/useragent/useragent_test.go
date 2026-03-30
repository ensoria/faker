package useragent_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/gofake/pkg/faker/common/util"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/generator/useragent"
	"github.com/ensoria/gofake/pkg/faker/provider"
	"github.com/ensoria/gofake/pkg/faker/provider/global"
	"github.com/ensoria/gofake/pkg/faker/testutil"
)

var _ = Describe("UserAgent", func() {
	coreRand := core.NewRand(util.RandSeed())
	g := &provider.Global{
		UserAgents: global.CreateUserAgents(),
	}
	ua := useragent.New(coreRand, g)

	Describe("RandomUserAgent", func() {
		It("should return a non-empty user agent string", func() {
			r := ua.RandomUserAgent()
			Expect(r).NotTo(BeEmpty())
			testutil.Output("UserAgent.RandomUserAgent", r)
		})
	})

	Describe("Chrome", func() {
		It("should return a Chrome user agent string", func() {
			r := ua.Chrome()
			Expect(r).To(HavePrefix("Mozilla/5.0 "))
			Expect(r).To(ContainSubstring("AppleWebKit/"))
			Expect(r).To(ContainSubstring("Chrome/"))
			Expect(r).To(ContainSubstring("Safari/"))
			testutil.Output("UserAgent.Chrome", r)
		})
	})

	Describe("MSEdge", func() {
		It("should return an Edge user agent string", func() {
			r := ua.MSEdge()
			Expect(r).To(HavePrefix("Mozilla/5.0 "))
			Expect(r).To(ContainSubstring("AppleWebKit/"))
			hasEdg := strings.Contains(r, "Edg/") || strings.Contains(r, "EdgA/") || strings.Contains(r, "EdgiOS/")
			// iOS Edge (EdgiOS) does not include "Chrome/" in the UA string, so only assert Chrome/ for non-iOS variants.
			if !strings.Contains(r, "EdgiOS/") {
				Expect(r).To(ContainSubstring("Chrome/"))
			}
			Expect(hasEdg).To(BeTrue())
			testutil.Output("UserAgent.MSEdge", r)
		})
	})

	Describe("Firefox", func() {
		It("should return a Firefox user agent string", func() {
			r := ua.Firefox()
			Expect(r).To(HavePrefix("Mozilla/5.0 "))
			Expect(r).To(ContainSubstring("Gecko/"))
			Expect(r).To(ContainSubstring("Firefox/"))
			testutil.Output("UserAgent.Firefox", r)
		})
	})

	Describe("Safari", func() {
		It("should return a Safari user agent string", func() {
			r := ua.Safari()
			Expect(r).To(HavePrefix("Mozilla/5.0 "))
			Expect(r).To(ContainSubstring("AppleWebKit/"))
			Expect(r).To(ContainSubstring("Safari/"))
			testutil.Output("UserAgent.Safari", r)
		})
	})

	Describe("Opera", func() {
		It("should return an Opera user agent string", func() {
			r := ua.Opera()
			Expect(r).To(HavePrefix("Opera/"))
			Expect(r).To(ContainSubstring("Presto/"))
			Expect(r).To(ContainSubstring("Version/"))
			testutil.Output("UserAgent.Opera", r)
		})
	})

	Describe("InternetExplorer", func() {
		It("should return an IE user agent string", func() {
			r := ua.InternetExplorer()
			Expect(r).To(HavePrefix("Mozilla/5.0 (compatible; MSIE "))
			Expect(r).To(ContainSubstring("Trident/"))
			testutil.Output("UserAgent.InternetExplorer", r)
		})
	})

	Describe("platform tokens", func() {
		It("WindowsPlatformToken should return a Windows token", func() {
			r := ua.WindowsPlatformToken()
			Expect(r).To(BeElementOf(g.UserAgents.WindowsPlatformTokens))
			testutil.Output("UserAgent.WindowsPlatformToken", r)
		})

		It("MacPlatformToken should return a Mac token", func() {
			r := ua.MacPlatformToken()
			Expect(r).To(HavePrefix("Macintosh; "))
			Expect(r).To(ContainSubstring("Mac OS X"))
			testutil.Output("UserAgent.MacPlatformToken", r)
		})

		It("IOSMobileToken should return an iOS token", func() {
			r := ua.IOSMobileToken()
			Expect(r).To(HavePrefix("iPhone; CPU iPhone OS"))
			Expect(r).To(ContainSubstring("like Mac OS X"))
			testutil.Output("UserAgent.IOSMobileToken", r)
		})

		It("LinuxPlatformToken should return a Linux token", func() {
			r := ua.LinuxPlatformToken()
			Expect(r).To(HavePrefix("X11; Linux "))
			testutil.Output("UserAgent.LinuxPlatformToken", r)
		})

		It("MacProcessor should return a Mac processor", func() {
			r := ua.MacProcessor()
			Expect(r).To(BeElementOf(g.UserAgents.MacProcessors))
			testutil.Output("UserAgent.MacProcessor", r)
		})

		It("LinuxProcessor should return a Linux processor", func() {
			r := ua.LinuxProcessor()
			Expect(r).To(BeElementOf(g.UserAgents.LinuxProcessors))
			testutil.Output("UserAgent.LinuxProcessor", r)
		})
	})
})
