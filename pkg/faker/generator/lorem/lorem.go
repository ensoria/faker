package lorem

import (
	"strings"

	"github.com/ensoria/gofake/pkg/faker/common/util"
	"github.com/ensoria/gofake/pkg/faker/core"
	"github.com/ensoria/gofake/pkg/faker/provider"
)

// Lorem provides methods for generating random lorem ipsum text.
//
// ランダムなLoremテキストを生成するメソッドを提供する構造体。
type Lorem struct {
	rand *core.Rand
	data *provider.Lorems
}

// New creates a new Lorem instance with the given random source and global data.
//
// 指定されたランダムソースとグローバルデータで新しいLoremインスタンスを作成する。
func New(rand *core.Rand, global *provider.Global) *Lorem {
	return &Lorem{
		rand,
		global.Lorems,
	}
}

// Word returns a random lorem word.
//
// ランダムなLoremの単語を返す。
func (l *Lorem) Word() string {
	return l.rand.Slice.StrElem(l.data.Words)
}

// WordSliceFixedLength returns a slice of random words with the specified count.
// num must be 1 or greater.
//
// 指定された数のランダムな単語のスライスを返す。
// numは1以上の整数。
func (l *Lorem) WordSliceFixedLength(num int) []string {
	if num < 1 {
		num = 1
	}
	var words []string
	for i := 0; i < num; i++ {
		words = append(words, l.Word())
	}
	return words
}

// WordSlice returns a slice of random words with a random count up to maxNum.
// maxNum must be 2 or greater.
//
// maxNum以下のランダムな数の単語のスライスを返す。
// maxNumは2以上の整数。
func (l *Lorem) WordSlice(maxNum int) []string {
	if maxNum < 2 {
		maxNum = 2
	}
	wordNum := l.rand.Num.IntBt(1, maxNum)
	return l.WordSliceFixedLength(wordNum)
}

// Words returns multiple random words as a single space-separated string.
//
// 複数のランダムな単語を1つのスペース区切りの文字列として返す。
func (l *Lorem) Words(num int) string {
	words := l.WordSlice(num)
	return strings.Join(words, " ")
}

// SentenceFixedLength returns a sentence with the specified number of words.
// The first letter is capitalized and the sentence ends with a period.
//
// 指定された単語数の文を返す。
// 最初の文字は大文字になり、末尾にピリオドが付く。
func (l *Lorem) SentenceFixedLength(wordNum int) string {
	words := l.WordSliceFixedLength(wordNum)
	words[0] = util.CapFirstLetter(words[0])
	return strings.Join(words, " ") + "."
}

// Sentence returns a sentence with a random number of words up to maxWordCount.
// maxWordCount must be 2 or greater.
//
// maxWordCount以下のランダムな単語数の文を返す。
// maxWordCountは2以上の整数。
func (l *Lorem) Sentence(maxWordCount int) string {
	if maxWordCount < 2 {
		maxWordCount = 2
	}
	wordNum := l.rand.Num.IntBt(1, maxWordCount)
	return l.SentenceFixedLength(wordNum)
}

// SentenceSliceFixedLength returns a slice of sentences with the specified count.
//
// 指定された数の文のスライスを返す。
func (l *Lorem) SentenceSliceFixedLength(sentenceNum int, wordMaxNum int) []string {
	var sentences []string
	for i := 0; i < sentenceNum; i++ {
		wordNum := l.rand.Num.IntBt(1, wordMaxNum)
		sentences = append(sentences, l.SentenceFixedLength(wordNum))
	}
	return sentences
}

// SentenceSlice returns a slice of sentences with a random count up to sentenceMaxNum.
//
// sentenceMaxNum以下のランダムな数の文のスライスを返す。
func (l *Lorem) SentenceSlice(sentenceMaxNum int, wordMaxNum int) []string {
	if sentenceMaxNum < 2 {
		sentenceMaxNum = 2
	}
	if wordMaxNum < 2 {
		wordMaxNum = 2
	}
	sentenceNum := l.rand.Num.IntBt(1, sentenceMaxNum)
	return l.SentenceSliceFixedLength(sentenceNum, wordMaxNum)
}

// Sentences returns multiple sentences as a single space-separated string.
//
// 複数の文を1つのスペース区切りの文字列として返す。
func (l *Lorem) Sentences(sentenceMaxNum int, wordMaxNum int) string {

	sentences := l.SentenceSlice(sentenceMaxNum, wordMaxNum)
	return strings.Join(sentences, " ")
}

// ParagraphSliceFixedLength returns a slice of paragraphs with the specified count.
// paragraphNum must be 1 or greater. sentenceMaxNum must be 2 or greater.
//
// 指定された数の段落のスライスを返す。
// paragraphNumは1以上、sentenceMaxNumは2以上の整数。
func (l *Lorem) ParagraphSliceFixedLength(paragraphNum int, sentenceMaxNum int) []string {
	if paragraphNum < 1 {
		paragraphNum = 1
	}
	if sentenceMaxNum < 2 {
		sentenceMaxNum = 2

	}
	var paragraphs []string
	for i := 0; i < paragraphNum; i++ {
		wordNum := l.rand.Num.IntBt(1, 30)
		sentenceNum := l.rand.Num.IntBt(1, sentenceMaxNum)
		paragraphs = append(paragraphs, l.Sentences(sentenceNum, wordNum))
	}
	return paragraphs
}

// ParagraphSlice returns a slice of paragraphs with a random count up to paragraphMaxNum.
//
// paragraphMaxNum以下のランダムな数の段落のスライスを返す。
func (l *Lorem) ParagraphSlice(paragraphMaxNum int, sentenceMaxNum int) []string {
	if paragraphMaxNum < 2 {
		paragraphMaxNum = 2
	}
	if sentenceMaxNum < 2 {
		sentenceMaxNum = 2
	}
	paragraphNum := l.rand.Num.IntBt(1, paragraphMaxNum)
	return l.ParagraphSliceFixedLength(paragraphNum, sentenceMaxNum)
}

// Paragraphs returns multiple paragraphs as a single newline-separated string.
//
// 複数の段落を1つの改行区切りの文字列として返す。
func (l *Lorem) Paragraphs(paragraphMaxNum int, sentenceMaxNum int) string {
	paragraphs := l.ParagraphSlice(paragraphMaxNum, sentenceMaxNum)
	return strings.Join(paragraphs, "\n\n")
}
