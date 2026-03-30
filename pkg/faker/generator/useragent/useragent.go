package useragent

import (
	"fmt"
	"time"

	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/provider"
)

const (
	mozillaPrefix = "Mozilla/5.0 "
)

// UserAgent provides methods for generating random user agent strings.
//
// ランダムなユーザーエージェント文字列を生成するメソッドを提供する構造体。
type UserAgent struct {
	rand *core.Rand
	data *provider.UserAgents
}

// New creates a new UserAgent instance with the given random source and global data.
//
// 指定されたランダムソースとグローバルデータで新しいUserAgentインスタンスを作成する。
func New(rand *core.Rand, global *provider.Global) *UserAgent {
	return &UserAgent{
		rand,
		global.UserAgents,
	}
}

// RandomUserAgent returns a random user agent string from any supported browser.
// Example: "Mozilla/5.0 (Windows CE) AppleWebKit/5350 (KHTML, like Gecko) Chrome/13.0.888.0 Safari/5350"
//
// サポートされているブラウザからランダムなユーザーエージェント文字列を返す。
func (u *UserAgent) RandomUserAgent() string {
	name := u.rand.Slice.StrElem(u.data.BrowserNames)

	switch name {
	case "firefox":
		return u.Firefox()
	case "chrome":
		return u.Chrome()
	case "internetExplorer":
		return u.InternetExplorer()
	case "opera":
		return u.Opera()
	case "safari":
		return u.Safari()
	case "msedge":
		return u.MSEdge()
	default:
		return u.Chrome()
	}
}

// Chrome returns a Chrome user agent string.
// Example: "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5) AppleWebKit/5312 (KHTML, like Gecko) Chrome/14.0.894.0 Safari/5312"
//
// Chromeのユーザーエージェント文字列を返す。
func (u *UserAgent) Chrome() string {
	saf := fmt.Sprintf("%d%d", u.rand.Num.IntBt(531, 536), u.rand.Num.IntBt(0, 2))
	chromeVer := fmt.Sprintf("%d.0.%d.0", u.rand.Num.IntBt(36, 40), u.rand.Num.IntBt(800, 899))

	platforms := []string{
		fmt.Sprintf("(%s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Mobile Safari/%s", u.LinuxPlatformToken(), saf, chromeVer, saf),
		fmt.Sprintf("(%s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Mobile Safari/%s", u.WindowsPlatformToken(), saf, chromeVer, saf),
		fmt.Sprintf("(%s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Mobile Safari/%s", u.MacPlatformToken(), saf, chromeVer, saf),
	}

	return mozillaPrefix + u.rand.Slice.StrElem(platforms)
}

// MSEdge returns a Microsoft Edge user agent string.
// Example: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.82 Safari/537.36 Edg/99.0.1150.36"
//
// Microsoft Edgeのユーザーエージェント文字列を返す。
func (u *UserAgent) MSEdge() string {
	saf := fmt.Sprintf("%d.%d", u.rand.Num.IntBt(531, 537), u.rand.Num.IntBt(0, 2))
	chrv := fmt.Sprintf("%d.0", u.rand.Num.IntBt(79, 99))
	chromeBuild := fmt.Sprintf(".%d.%d", u.rand.Num.IntBt(4000, 4844), u.rand.Num.IntBt(10, 99))
	edgeBuild := fmt.Sprintf("%d.%d", u.rand.Num.IntBt(1000, 1146), u.rand.Num.IntBt(0, 99))

	platforms := []string{
		fmt.Sprintf("(%s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s%s Safari/%s Edg/%s%s", u.WindowsPlatformToken(), saf, chrv, chromeBuild, saf, chrv, edgeBuild),
		fmt.Sprintf("(%s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s%s Safari/%s Edg/%s%s", u.MacPlatformToken(), saf, chrv, chromeBuild, saf, chrv, edgeBuild),
		fmt.Sprintf("(%s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s%s Safari/%s EdgA/%s%s", u.LinuxPlatformToken(), saf, chrv, chromeBuild, saf, chrv, edgeBuild),
		fmt.Sprintf("(%s) AppleWebKit/%s (KHTML, like Gecko) Version/15.0 EdgiOS/%s%s Mobile/15E148 Safari/%s", u.IOSMobileToken(), saf, chrv, edgeBuild, saf),
	}

	return mozillaPrefix + u.rand.Slice.StrElem(platforms)
}

// Firefox returns a Firefox user agent string.
// Example: "Mozilla/5.0 (X11; Linux i686; rv:7.0) Gecko/20101231 Firefox/3.6"
//
// Firefoxのユーザーエージェント文字列を返す。
func (u *UserAgent) Firefox() string {
	// Random date between 2010-01-01 and now
	from := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)
	randDate := u.rand.Time.TimeRange(from, time.Now())
	geckoDate := randDate.Format("20060102")
	ver := fmt.Sprintf("Gecko/%s Firefox/%d.0", geckoDate, u.rand.Num.IntBt(35, 37))

	lang := u.rand.Slice.StrElem(u.data.Languages)

	platforms := []string{
		fmt.Sprintf("(%s; %s; rv:1.9.%d.20) %s", u.WindowsPlatformToken(), lang, u.rand.Num.IntBt(0, 2), ver),
		fmt.Sprintf("(%s; rv:%d.0) %s", u.LinuxPlatformToken(), u.rand.Num.IntBt(5, 7), ver),
		fmt.Sprintf("(%s rv:%d.0) %s", u.MacPlatformToken(), u.rand.Num.IntBt(2, 6), ver),
	}

	return mozillaPrefix + u.rand.Slice.StrElem(platforms)
}

// Safari returns a Safari user agent string.
// Example: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_7_1 rv:3.0; en-US) AppleWebKit/534.11.3 (KHTML, like Gecko) Version/4.0 Safari/534.11.3"
//
// Safariのユーザーエージェント文字列を返す。
func (u *UserAgent) Safari() string {
	saf := fmt.Sprintf("%d.%d.%d", u.rand.Num.IntBt(531, 535), u.rand.Num.IntBt(1, 50), u.rand.Num.IntBt(1, 7))

	var ver string
	if u.rand.Bool.Evenly() {
		ver = fmt.Sprintf("%d.%d", u.rand.Num.IntBt(4, 5), u.rand.Num.IntBt(0, 1))
	} else {
		ver = fmt.Sprintf("%d.0.%d", u.rand.Num.IntBt(4, 5), u.rand.Num.IntBt(1, 5))
	}

	mobileDevices := []string{
		"iPhone; CPU iPhone OS",
		"iPad; CPU OS",
	}

	lang := u.rand.Slice.StrElem(u.data.Languages)

	platforms := []string{
		fmt.Sprintf("(Windows; U; %s) AppleWebKit/%s (KHTML, like Gecko) Version/%s Safari/%s", u.WindowsPlatformToken(), saf, ver, saf),
		fmt.Sprintf("(%s rv:%d.0; %s) AppleWebKit/%s (KHTML, like Gecko) Version/%s Safari/%s", u.MacPlatformToken(), u.rand.Num.IntBt(2, 6), lang, saf, ver, saf),
		fmt.Sprintf("(%s %d_%d_%d like Mac OS X; %s) AppleWebKit/%s (KHTML, like Gecko) Version/%d.0.5 Mobile/8B%d Safari/6%s",
			u.rand.Slice.StrElem(mobileDevices),
			u.rand.Num.IntBt(7, 8), u.rand.Num.IntBt(0, 2), u.rand.Num.IntBt(1, 2),
			lang, saf, u.rand.Num.IntBt(3, 4), u.rand.Num.IntBt(111, 119), saf),
	}

	return mozillaPrefix + u.rand.Slice.StrElem(platforms)
}

// Opera returns an Opera user agent string.
// Example: "Opera/8.25 (Windows NT 5.1; en-US) Presto/2.9.188 Version/10.00"
//
// Operaのユーザーエージェント文字列を返す。
func (u *UserAgent) Opera() string {
	lang := u.rand.Slice.StrElem(u.data.Languages)
	presto := fmt.Sprintf("Presto/2.%d.%d Version/%d.00", u.rand.Num.IntBt(8, 12), u.rand.Num.IntBt(160, 355), u.rand.Num.IntBt(10, 12))

	platforms := []string{
		fmt.Sprintf("(%s; %s) %s", u.LinuxPlatformToken(), lang, presto),
		fmt.Sprintf("(%s; %s) %s", u.WindowsPlatformToken(), lang, presto),
	}

	return fmt.Sprintf("Opera/%d.%d %s", u.rand.Num.IntBt(8, 9), u.rand.Num.IntBt(10, 99), u.rand.Slice.StrElem(platforms))
}

// InternetExplorer returns an Internet Explorer user agent string.
// Example: "Mozilla/5.0 (compatible; MSIE 7.0; Windows 98; Win 9x 4.90; Trident/3.0)"
//
// Internet Explorerのユーザーエージェント文字列を返す。
func (u *UserAgent) InternetExplorer() string {
	return fmt.Sprintf("Mozilla/5.0 (compatible; MSIE %d.0; %s; Trident/%d.%d)",
		u.rand.Num.IntBt(5, 11), u.WindowsPlatformToken(), u.rand.Num.IntBt(3, 5), u.rand.Num.IntBt(0, 1))
}

// WindowsPlatformToken returns a random Windows platform token.
// Example: "Windows NT 6.1"
//
// ランダムなWindowsプラットフォームトークンを返す。
func (u *UserAgent) WindowsPlatformToken() string {
	return u.rand.Slice.StrElem(u.data.WindowsPlatformTokens)
}

// MacPlatformToken returns a random Mac platform token.
// Example: "Macintosh; Intel Mac OS X 10_7_3"
//
// ランダムなMacプラットフォームトークンを返す。
func (u *UserAgent) MacPlatformToken() string {
	return fmt.Sprintf("Macintosh; %s Mac OS X 10_%d_%d",
		u.rand.Slice.StrElem(u.data.MacProcessors), u.rand.Num.IntBt(5, 8), u.rand.Num.IntBt(0, 9))
}

// IOSMobileToken returns a random iOS mobile platform token.
// Example: "iPhone; CPU iPhone OS 14_1 like Mac OS X"
//
// ランダムなiOSモバイルプラットフォームトークンを返す。
func (u *UserAgent) IOSMobileToken() string {
	return fmt.Sprintf("iPhone; CPU iPhone OS %d_%d like Mac OS X",
		u.rand.Num.IntBt(13, 15), u.rand.Num.IntBt(0, 2))
}

// LinuxPlatformToken returns a random Linux platform token.
// Example: "X11; Linux x86_64"
//
// ランダムなLinuxプラットフォームトークンを返す。
func (u *UserAgent) LinuxPlatformToken() string {
	return fmt.Sprintf("X11; Linux %s", u.rand.Slice.StrElem(u.data.LinuxProcessors))
}

// MacProcessor returns a random Mac processor string.
// Example: "Intel"
//
// ランダムなMacプロセッサ文字列を返す。
func (u *UserAgent) MacProcessor() string {
	return u.rand.Slice.StrElem(u.data.MacProcessors)
}

// LinuxProcessor returns a random Linux processor string.
// Example: "x86_64"
//
// ランダムなLinuxプロセッサ文字列を返す。
func (u *UserAgent) LinuxProcessor() string {
	return u.rand.Slice.StrElem(u.data.LinuxProcessors)
}
