package file

import (
	"io"
	"os"
	"path/filepath"

	"github.com/ensoria/faker/pkg/faker/core"
	"github.com/ensoria/faker/pkg/faker/provider"
)

// File provides methods for generating random file-related data and creating files.
//
// ランダムなファイル関連データの生成やファイル作成のメソッドを提供する構造体。
type File struct {
	rand *core.Rand
	data *provider.Files
}

// New creates a new File instance with the given random source and global data.
//
// 指定されたランダムソースとグローバルデータで新しいFileインスタンスを作成する。
func New(rand *core.Rand, global *provider.Global) *File {
	return &File{
		rand,
		global.Files,
	}
}

func (f *File) randomEntry() *provider.MIMEEntry {
	return f.data.MIMEEntries[f.rand.Num.Intn(len(f.data.MIMEEntries))]
}

// MIMEType returns a random MIME type string.
//
// ランダムなMIMEタイプ文字列を返す。
func (f *File) MIMEType() string {
	return f.randomEntry().Type
}

// Extension returns a random file extension.
//
// ランダムなファイル拡張子を返す。
func (f *File) Extension() string {
	entry := f.randomEntry()
	return f.rand.Slice.StrElem(entry.Extensions)
}

// WriteWithText creates a file with the given text content and extension.
// The file name is randomly generated. The extension should not include a dot.
// If returnFullPath is true, the absolute path is returned.
//
// 指定されたテキスト内容と拡張子でファイルを作成する。
// ファイル名はランダムに生成される。拡張子にドットは含めない。
// returnFullPathがtrueの場合、絶対パスを返す。
func (f *File) WriteWithText(
	destDir string,
	content string,
	extension string,
	returnFullPath bool,
) (string, error) {

	DirErr := os.MkdirAll(destDir, 0777)
	if DirErr != nil {
		return "", DirErr
	}

	// 16 letters
	fileName := f.rand.Str.AlphaDigitsLike("****************") + "." + extension
	filePath := filepath.Join(destDir, fileName)

	data := []byte(content)
	FileErr := os.WriteFile(filePath, data, 0777)
	if FileErr != nil {
		return "", FileErr
	}

	if returnFullPath {
		return filepath.Abs(filePath)
	}

	return filePath, nil
}

// CopyFrom creates a file by copying content from a source file.
// The file name is randomly generated. The extension should not include a dot.
// If returnFullPath is true, the absolute path is returned.
//
// ソースファイルからコンテンツをコピーしてファイルを作成する。
// ファイル名はランダムに生成される。拡張子にドットは含めない。
// returnFullPathがtrueの場合、絶対パスを返す。
func (f *File) CopyFrom(
	destDir string,
	srcFilePath string,
	extension string,
	returnFullPath bool,
) (string, error) {
	srcFile, openErr := os.Open(srcFilePath)
	if openErr != nil {
		return "", openErr
	}
	defer srcFile.Close()

	DirErr := os.MkdirAll(destDir, 0777)
	if DirErr != nil {
		return "", DirErr
	}
	// 16 letters
	fileName := f.rand.Str.AlphaDigitsLike("****************") + "." + extension
	filePath := filepath.Join(destDir, fileName)
	destFile, createErr := os.Create(filePath)
	if createErr != nil {
		return "", createErr
	}
	defer destFile.Close()

	_, copyErr := io.Copy(destFile, srcFile)
	if copyErr != nil {
		return "", copyErr
	}

	if returnFullPath {
		return filepath.Abs(filePath)
	}

	return filePath, nil
}
