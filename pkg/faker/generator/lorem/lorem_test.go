package lorem_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/lorem"
	"github.com/ensoria/faker/pkg/faker/provider"
	"github.com/ensoria/faker/pkg/faker/provider/global"
	"github.com/ensoria/faker/pkg/faker/testutil"
)

var _ = Describe("Lorem", func() {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Lorems: global.CreateLorems(),
	}

	lm := lorem.New(coreRand, global)

	It("Word should return a word", func() {
		r := lm.Word()

		testutil.Output("Lorem.Word", r)
	})

	It("WordSliceFixedLength should return a slice of words with fixed length", func() {
		r := lm.WordSliceFixedLength(5)
		ln := len(r)
		Expect(ln).To(BeNumerically("==", 5))
	})

	It("WordSlice should return a slice of words less than specified length", func() {
		r := lm.WordSlice(5)

		Expect(len(r)).To(BeNumerically(">", 0))
		Expect(len(r)).To(BeNumerically("<=", 5))

		testutil.Output("Lorem.WordSlice", r)
	})

	It("Words should return a string of words less than specified number", func() {
		r := lm.Words(5)

		words := strings.Split(strings.TrimSpace(r), " ")
		Expect(len(words)).To(BeNumerically(">", 0))
		Expect(len(words)).To(BeNumerically("<=", 5))
		testutil.Output("Lorem.Words", r)
	})

	It("SentenceFixedLength should return a sentence with fixed length", func() {
		r := lm.SentenceFixedLength(6)

		ln := len(strings.Split(r, " "))
		Expect(ln).To(BeNumerically("==", 6))
		testutil.Output("Lorem.SentenceFixedLength", r)
	})

	It("Sentence should return a sentence", func() {
		r := lm.Sentence(6)

		ln := len(strings.Split(r, " "))
		Expect(ln).To(BeNumerically(">", 0))
		Expect(ln).To(BeNumerically("<=", 6))
		testutil.Output("Lorem.Sentence", r)
	})

	It("SentenceSliceFixedLength should return a slice of sentences with fixed length", func() {
		r := lm.SentenceSliceFixedLength(3, 6)

		ln := len(r)
		Expect(ln).To(BeNumerically("==", 3))
		for i, v := range r {
			wordCount := len(strings.Split(v, " "))
			Expect(wordCount).To(BeNumerically(">", 0))
			Expect(wordCount).To(BeNumerically("<=", 6))
			o := fmt.Sprintf("Lorem.SentenceSliceFixedLength(%d)", i+1)
			testutil.Output(o, v)
		}
	})

	It("SentenceSlice should return slices of sentences", func() {
		r := lm.SentenceSlice(3, 6)

		l := len(r)
		Expect(l).To(BeNumerically(">", 0))
		Expect(l).To(BeNumerically("<=", 3))
		for i, v := range r {
			wordCount := len(strings.Split(v, " "))
			Expect(wordCount).To(BeNumerically(">", 0))
			Expect(wordCount).To(BeNumerically("<=", 6))
			o := fmt.Sprintf("Lorem.SentenceSlice(%d)", i+1)
			testutil.Output(o, v)
		}
	})

	It("Sentences should return a string of sentences", func() {
		r := lm.Sentences(3, 6)

		sentences := strings.Split(strings.TrimSpace(r), ". ")
		Expect(len(sentences)).To(BeNumerically(">", 0))
		Expect(len(sentences)).To(BeNumerically("<=", 3))
		for i, v := range sentences {
			wordCount := len(strings.Split(v, " "))
			Expect(wordCount).To(BeNumerically(">", 0))
			Expect(wordCount).To(BeNumerically("<=", 6))
			o := fmt.Sprintf("Lorem.Sentences(%d)", i+1)
			testutil.Output(o, v)
		}
	})

	It("ParagraphSliceFixedLength should return a paragraph with fixed length of sentences", func() {
		r := lm.ParagraphSliceFixedLength(3, 6)

		l := len(r)
		Expect(l).To(BeNumerically("==", 3))
		for i, v := range r {
			sentences := strings.Split(strings.TrimSpace(v), ". ")
			Expect(len(sentences)).To(BeNumerically(">", 0))
			Expect(len(sentences)).To(BeNumerically("<=", 6))
			o := fmt.Sprintf("Lorem.ParagraphSliceFixedLength(%d)", i+1)
			testutil.Output(o, v)
		}
	})

	It("ParagraphSlice should return a paragraph with random length of sentences", func() {
		r := lm.ParagraphSlice(3, 6)

		l := len(r)
		Expect(l).To(BeNumerically(">", 0))
		Expect(l).To(BeNumerically("<=", 3))
		for i, v := range r {
			sentences := strings.Split(strings.TrimSpace(v), ". ")
			Expect(len(sentences)).To(BeNumerically(">", 0))
			Expect(len(sentences)).To(BeNumerically("<=", 6))
			o := fmt.Sprintf("Lorem.ParagraphSlice(%d)", i+1)
			testutil.Output(o, v)
		}
	})

	It("Paragraphs should return a string of paragraphs", func() {
		r := lm.Paragraphs(3, 6)

		paragraphs := strings.Split(strings.TrimSpace(r), "\n\n")
		Expect(len(paragraphs)).To(BeNumerically(">", 0))
		Expect(len(paragraphs)).To(BeNumerically("<=", 3))
		for i, v := range paragraphs {
			sentences := strings.Split(strings.TrimSpace(v), ". ")
			Expect(len(sentences)).To(BeNumerically(">", 0))
			Expect(len(sentences)).To(BeNumerically("<=", 6))
			o := fmt.Sprintf("Lorem.Paragraphs(%d)", i+1)
			testutil.Output(o, v)
		}
	})

})
