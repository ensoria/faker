package log

import (
	"log"
	"runtime"
)

const (
	unavailableLocaleMsg = "This method is not available for the current locale."
	wrongUsageMsgFormat  = `faker might-be-wrong-usage warning. Empty value of the type returned: "%s" at [%s]: line [%d]: called: [%s]`
	generalErrorFormat   = "%s at [%s]: line [%d]: called: [%s]"
)

// logger is nil by default, meaning no log output.
// Use SetLogger to enable logging.
//
// loggerはデフォルトでnilであり、ログは出力されない。
// ログ出力を有効にするにはSetLoggerを使用する。
var logger *log.Logger

// SetLogger sets the logger used by the faker package.
// Pass nil to disable logging (default).
//
// fakerパッケージが使用するloggerを設定する。
// nilを渡すとログ出力が無効になる（デフォルト）。
func SetLogger(l *log.Logger) {
	logger = l
}

// GetCallerInfo returns the function, file path, and line number of the caller at the given skip level.
//
// 指定されたスキップレベルの呼び出し元の関数、ファイルパス、行番号を返す。
func GetCallerInfo(skip int) (*runtime.Func, string, int) {
	trueSkip := skip + 1
	_, file, line, _ := runtime.Caller(trueSkip)
	pc, _, _, _ := runtime.Caller(skip) // the function/method itself
	called := runtime.FuncForPC(pc)
	return called, file, line
}

// UnavailableLocale logs a warning that the called method is not available for the current locale.
//
// 呼び出されたメソッドが現在のロケールで利用できないことを警告としてログに出力する。
func UnavailableLocale(skip int) {
	WrongUsage(unavailableLocaleMsg, skip+1)
}

// WrongUsage logs a warning about potentially incorrect usage of the faker library,
// including caller information.
//
// fakerライブラリの不正な使用の可能性について、呼び出し元の情報を含めて警告をログに出力する。
func WrongUsage(msg string, skip int) {
	if logger == nil {
		return
	}
	trueSkip := skip + 1
	caller, file, line := GetCallerInfo(trueSkip)
	logger.Printf(wrongUsageMsgFormat, msg, file, line, caller.Name())
}

// GeneralError logs a general error message with caller information.
//
// 呼び出し元の情報を含めて一般的なエラーメッセージをログに出力する。
func GeneralError(msg string, skip int) {
	if logger == nil {
		return
	}
	trueSkip := skip + 1
	caller, file, line := GetCallerInfo(trueSkip)
	logger.Printf(generalErrorFormat, msg, file, line, caller.Name())
}
