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
var logger *log.Logger

// SetLogger sets the logger used by the faker package.
// Pass nil to disable logging (default).
func SetLogger(l *log.Logger) {
	logger = l
}

func GetCallerInfo(skip int) (*runtime.Func, string, int) {
	trueSkip := skip + 1
	_, file, line, _ := runtime.Caller(trueSkip)
	pc, _, _, _ := runtime.Caller(skip) // the function/method itself
	called := runtime.FuncForPC(pc)
	return called, file, line
}

func UnavailableLocale(skip int) {
	WrongUsage(unavailableLocaleMsg, skip+1)
}

func WrongUsage(msg string, skip int) {
	if logger == nil {
		return
	}
	trueSkip := skip + 1
	caller, file, line := GetCallerInfo(trueSkip)
	logger.Printf(wrongUsageMsgFormat, msg, file, line, caller.Name())
}

func GeneralError(msg string, skip int) {
	if logger == nil {
		return
	}
	trueSkip := skip + 1
	caller, file, line := GetCallerInfo(trueSkip)
	logger.Printf(generalErrorFormat, msg, file, line, caller.Name())
}
