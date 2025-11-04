package file_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ensoria/faker/pkg/faker/common/util"
	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/generator/file"
	"github.com/ensoria/faker/pkg/faker/provider"
	"github.com/ensoria/faker/pkg/faker/provider/global"
)

var _ = Describe("File", Ordered, func() {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Files: global.CreateFiles(),
	}

	f := file.New(coreRand, global)

	Describe("File", func() {
		It("MimeType should return a string", func() {
			mimeType := f.MimeType()
			_, exists := global.Files.MimeTypes[mimeType]
			Expect(exists).To(BeTrue())
		})

		It("Extension should return a string", func() {
			extension := f.Extension()
			Expect(extension).To(MatchRegexp(`^[0-9a-z-]{2,20}$`))
		})
	})
	// NOTICE: you could use `t.TempDir() to create a temporary directory.
	// But sometimes I want to see what's in the directory for debugging.
	// So I'm using a fixed directory here.
	destDir := "../../../../tmp"

	Describe("File", func() {
		content := "Hello, World!"
		extension := "txt"

		When("returnFullPath false", func() {
			returnFullPath := false
			It("WriteWithText should create file and return a relative path", func() {
				filePath, err := f.WriteWithText(destDir, content, extension, returnFullPath)
				r := FileExists(filePath)
				Expect(r).To(BeTrue())
				Expect(filePath).To(MatchRegexp(`^..\/..\/..\/..\/tmp/[\d\w]{16}\.txt$`))
				Expect(err).To(BeNil())
			})
		})

		When("returnFullPath true", func() {
			returnFullPath := true
			It("WriteWithText should return absolute path", func() {
				filePath, err := f.WriteWithText(destDir, content, extension, returnFullPath)
				r := FileExists(filePath)
				Expect(r).To(BeTrue())
				absPath, _ := filepath.Abs(filePath)
				Expect(filePath).To(Equal(absPath))
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("CopyFrom", func() {
		srcFilePath := "../../../../testdata/files/sample.txt"
		extension := "txt"

		When("returnFullPath is false", func() {
			returnFullPath := false
			It("CopyFrom should copy file from src to dest and should return relative path", func() {
				filePath, err := f.CopyFrom(destDir, srcFilePath, extension, returnFullPath)
				r := FileExists(filePath)
				Expect(r).To(BeTrue())
				Expect(filePath).To(MatchRegexp(`^..\/..\/..\/..\/tmp/[\d\w]{16}\.txt$`))
				Expect(err).To(BeNil())
			})
		})

		When("returnFullPath is true", func() {
			returnFullPath := true
			It("CopyFrom should copy file from src to dest and should return full path", func() {
				filePath, err := f.CopyFrom(destDir, srcFilePath, extension, returnFullPath)
				r := FileExists(filePath)
				Expect(r).To(BeTrue())
				absPath, _ := filepath.Abs(filePath)
				Expect(filePath).To(Equal(absPath))
				Expect(err).To(BeNil())
			})
		})
	})

	AfterAll(func() {
		// Clean up
		deleteFilesInDir(destDir)
	})
})

func FileExists(filePath string) bool {
	_, pathErr := os.Stat(filePath)
	r := os.IsNotExist(pathErr) // DO NOT USE IsExist().
	return !r
}

func deleteFilesInDir(dirPath string) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		err = os.Remove(filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
