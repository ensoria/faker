package person

import (
	"github.com/ensoria/gofake/pkg/faker/common/log"
	"github.com/ensoria/gofake/pkg/faker/common/util"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/provider"
)

// Person provides methods for generating random person data.
//
// ランダムな人物データを生成するメソッドを提供する構造体。
type Person struct {
	rand *core.Rand
	data *provider.People
}

// New creates a new Person instance with the given random source and locale data.
//
// 指定されたランダムソースとロケールデータで新しいPersonインスタンスを作成する。
func New(rand *core.Rand, localized *provider.Localized) *Person {
	return &Person{
		rand,
		localized.People,
	}
}

// FirstNameMale returns a random male first name.
//
// ランダムな男性のファーストネームを返す。
func (p *Person) FirstNameMale() string {
	if len(p.data.FirstNameMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.FirstNameMales)
}

// FirstNameFemale returns a random female first name.
//
// ランダムな女性のファーストネームを返す。
func (p *Person) FirstNameFemale() string {
	if len(p.data.FirstNameFemales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.FirstNameFemales)
}

// FirstName returns a random first name (male or female with equal probability).
//
// ランダムなファーストネームを返す（男女等確率）。
func (p *Person) FirstName() string {
	if len(p.data.FirstNameFemales) == 0 || len(p.data.FirstNameMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	if p.rand.Bool.Evenly() {
		return p.rand.Slice.StrElem(p.data.FirstNameMales)
	}
	return p.rand.Slice.StrElem(p.data.FirstNameFemales)
}

// LastName returns a random last name.
//
// ランダムなラストネームを返す。
func (p *Person) LastName() string {
	if len(p.data.LastNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	return p.rand.Slice.StrElem(p.data.LastNames)
}

// TitleMale returns a random male title.
//
// ランダムな男性の敬称を返す。
func (p *Person) TitleMale() string {
	if len(p.data.TitleMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.TitleMales)
}

// TitleFemale returns a random female title.
//
// ランダムな女性の敬称を返す。
func (p *Person) TitleFemale() string {
	if len(p.data.TitleFemales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.TitleFemales)
}

// Title returns a random title (male or female with equal probability).
//
// ランダムな敬称を返す（男女等確率）。
func (p *Person) Title() string {
	if len(p.data.TitleFemales) == 0 || len(p.data.TitleMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	if p.rand.Bool.Evenly() {
		return p.rand.Slice.StrElem(p.data.TitleMales)
	}
	return p.rand.Slice.StrElem(p.data.TitleFemales)
}

// Suffix returns a random name suffix (e.g. Jr., PhD).
//
// ランダムな名前の接尾辞を返す（例: Jr., PhD）。
func (p *Person) Suffix() string {
	if len(p.data.Suffixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	return p.rand.Slice.StrElem(p.data.Suffixes)
}

// MaleName returns a randomly formatted male full name using locale-specific formats.
//
// ロケール固有のフォーマットを使ってランダムにフォーマットされた男性のフルネームを返す。
func (p *Person) MaleName() string {
	if len(p.data.MaleNameFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := p.rand.Slice.StrElem(p.data.MaleNameFormats)
	nameData := p.data.CreateNameMale(p)
	return util.RenderTemplate(format, nameData)
}

// FemaleName returns a randomly formatted female full name using locale-specific formats.
//
// ロケール固有のフォーマットを使ってランダムにフォーマットされた女性のフルネームを返す。
func (p *Person) FemaleName() string {
	if len(p.data.FemaleNameFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := p.rand.Slice.StrElem(p.data.FemaleNameFormats)
	nameData := p.data.CreateNameFemale(p)
	return util.RenderTemplate(format, nameData)
}

// Name returns a randomly formatted full name (male or female with equal probability).
//
// ランダムにフォーマットされたフルネームを返す（男女等確率）。
func (p *Person) Name() string {
	if (len(p.data.MaleNameFormats) == 0) || (len(p.data.FemaleNameFormats) == 0) {
		log.UnavailableLocale(1)
		return ""
	}

	if p.rand.Bool.Evenly() {
		return p.MaleName()
	}
	return p.FemaleName()
}

// SSN returns a random Social Security Number.
//
// ランダムな社会保障番号を返す。
func (p *Person) SSN() string {
	return p.rand.Str.AlphaDigitsLike("###-##-####")
}

// FirstKanaNameMale returns a random male first name in katakana (for ja_JP locale).
//
// ランダムな男性のカタカナのファーストネームを返す（ja_JPロケール用）。
func (p *Person) FirstKanaNameMale() string {
	if len(p.data.FirstKanaNameMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.FirstKanaNameMales)
}

// FirstKanaNameFemale returns a random female first name in katakana (for ja_JP locale).
//
// ランダムな女性のカタカナのファーストネームを返す（ja_JPロケール用）。
func (p *Person) FirstKanaNameFemale() string {
	if len(p.data.FirstKanaNameFemales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.FirstKanaNameFemales)
}

// FirstKanaName returns a random first name in katakana (male or female with equal probability).
//
// ランダムなカタカナのファーストネームを返す（男女等確率）。
func (p *Person) FirstKanaName() string {
	if len(p.data.FirstKanaNameFemales) == 0 || len(p.data.FirstKanaNameMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	if p.rand.Bool.Evenly() {
		return p.rand.Slice.StrElem(p.data.FirstKanaNameMales)
	}
	return p.rand.Slice.StrElem(p.data.FirstKanaNameFemales)
}

// LastKanaName returns a random last name in katakana (for ja_JP locale).
//
// ランダムなカタカナのラストネームを返す（ja_JPロケール用）。
func (p *Person) LastKanaName() string {
	if len(p.data.LastKanaNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	return p.rand.Slice.StrElem(p.data.LastKanaNames)
}

// MaleKanaName returns a randomly formatted male full name in katakana.
//
// ランダムにフォーマットされた男性のカタカナのフルネームを返す。
func (p *Person) MaleKanaName() string {
	if len(p.data.MaleNameFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := p.rand.Slice.StrElem(p.data.MaleNameFormats)
	nameData := p.data.CreateKanaNameMale(p)
	return util.RenderTemplate(format, nameData)
}

// FemaleKanaName returns a randomly formatted female full name in katakana.
//
// ランダムにフォーマットされた女性のカタカナのフルネームを返す。
func (p *Person) FemaleKanaName() string {
	if len(p.data.FemaleNameFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := p.rand.Slice.StrElem(p.data.FemaleNameFormats)
	nameData := p.data.CreateKanaNameFemale(p)
	return util.RenderTemplate(format, nameData)
}

// KanaName returns a randomly formatted full name in katakana (male or female with equal probability).
//
// ランダムにフォーマットされたカタカナのフルネームを返す（男女等確率）。
func (p *Person) KanaName() string {
	if (len(p.data.MaleNameFormats) == 0) || (len(p.data.FemaleNameFormats) == 0) {
		log.UnavailableLocale(1)
		return ""
	}

	if p.rand.Bool.Evenly() {
		return p.MaleKanaName()
	}
	return p.FemaleKanaName()
}
