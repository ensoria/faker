package internet

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/provider"
)

const (
	defaultSlugWordCount   = 6
	slugVarianceMin        = 60
	slugVarianceMax        = 140
	slugVarianceDivisor    = 100
	passwordMinLength      = 8
	passwordMaxLength      = 20
	passwordAlphaDigitChar = '*'
	passwordSpecialChar    = '!'
)

type Internet struct {
	rand  *core.Rand
	data  *provider.Internets
	words []string
}

func New(rand *core.Rand, global *provider.Global) *Internet {
	return &Internet{
		rand:  rand,
		data:  global.Internets,
		words: global.Lorems.Words,
	}
}

func (i *Internet) FirstName() string {
	fName := i.rand.Slice.StrElem(i.data.FirstNames)
	without1Quote := strings.ReplaceAll(fName, "'", "")
	return strings.ToLower(without1Quote)
}

func (i *Internet) LastName() string {
	lName := i.rand.Slice.StrElem(i.data.LastNames)
	without1Quote := strings.ReplaceAll(lName, "'", "")
	return strings.ToLower(without1Quote)
}

func (i *Internet) UserName() string {
	baseFormat := i.rand.Slice.StrElem(i.data.UserNameFormats)
	format := i.rand.Str.AlphaDigitsLike(baseFormat)
	userName := i.data.CreateUserName(i)
	return util.RenderTemplate(format, userName)
}

func (i *Internet) DomainWord() string {
	lastName := i.rand.Slice.StrElem(i.data.LastNames)
	word := strings.ToLower(lastName)
	return word
}

func (i *Internet) TLD() string {
	return i.rand.Slice.StrElem(i.data.TLD)
}

// example: "howell.com"
func (i *Internet) DomainName() string {
	return i.DomainWord() + "." + i.TLD()
}

// example: "jude.borer@oberbrunner.com"
func (i *Internet) Email() string {
	format := i.rand.Slice.StrElem(i.data.EmailFormats)
	data := i.data.CreateEmail(i)
	return util.RenderTemplate(format, data)
}

// Password generates a random password string.
// If includeSpecial is true, some characters will be special characters.
// Example: "18w50q2412G5Iky60QL" (without special)
// Example: "k3$Rp8!mZ2&xQ" (with special)
func (i *Internet) Password(includeSpecial bool) string {
	length := i.rand.Num.IntBt(passwordMinLength, passwordMaxLength)
	var like []byte
	for idx := 0; idx < length; idx++ {
		if includeSpecial && i.rand.Num.Intn(4) == 0 {
			like = append(like, byte(passwordSpecialChar))
		} else {
			like = append(like, byte(passwordAlphaDigitChar))
		}
	}
	return i.rand.Str.AlphaDigitsLike(string(like))
}

// Slug generates a URL slug from random lorem words joined by hyphens.
// If variableWordCount is true, the actual word count varies around nbWords.
// Example: "aut-repellat-commodi-vel-itaque-nihil"
func (i *Internet) Slug(nbWords int, variableWordCount bool) string {
	if nbWords <= 0 {
		nbWords = defaultSlugWordCount
	}

	if variableWordCount {
		nbWords = nbWords*i.rand.Num.IntBt(slugVarianceMin, slugVarianceMax)/slugVarianceDivisor + 1
	}

	words := make([]string, nbWords)
	for idx := 0; idx < nbWords; idx++ {
		words[idx] = i.rand.Slice.StrElem(i.words)
	}

	return strings.Join(words, "-")
}

// URL generates a random URL.
// Example: "http://www.runolfsdottir.com/aut-repellat-commodi"
func (i *Internet) URL() string {
	format := i.rand.Slice.StrElem(i.data.URLFormats)

	data := &urlData{
		DomainName: i.DomainName(),
		Slug:       i.Slug(defaultSlugWordCount, true),
	}

	return util.RenderTemplate(format, data)
}

type urlData struct {
	DomainName string
	Slug       string
}

func (i *Internet) IPv4() net.IP {
	var ipNum int
	if i.rand.Bool.Evenly() {
		ipNum = i.rand.Num.IntBt(-2147483648, -2)
	} else {
		ipNum = i.rand.Num.IntBt(16777216, 2147483647)
	}

	return uint32ToIP(uint32(ipNum))
}

func (i *Internet) IPv6() string {
	var res []string

	for index := 0; index < 8; index++ {
		res = append(res, fmt.Sprintf("%x", i.rand.Num.Intn(65536)))
	}

	return strings.Join(res, ":")

}

func (i *Internet) LocalIPv4() net.IP {
	lenIPBlocks := len(i.data.LocalIPBlocks)
	ipBlock := i.data.LocalIPBlocks[i.rand.Num.Intn(lenIPBlocks)]
	ipBlock0, _ := ipToUint32(ipBlock[0])
	ipBlock1, _ := ipToUint32(ipBlock[1])
	num := i.rand.Num.Int32Bt(int32(ipBlock0), int32(ipBlock1))
	return uint32ToIP(uint32(num))
}

func (i *Internet) MACAddress() string {
	var mac []string

	for index := 0; index < 6; index++ {
		mac = append(mac, fmt.Sprintf("%02X", i.rand.Num.Intn(256)))
	}

	return strings.Join(mac, ":")
}

func uint32ToIP(long uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, long)
	return ip
}

func ipToUint32(ipStr string) (uint32, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0, fmt.Errorf("invalid IP address: %s", ipStr)
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}
