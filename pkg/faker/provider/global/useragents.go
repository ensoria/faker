package global

import "github.com/ensoria/gofake/pkg/faker/provider"

func CreateUserAgents() *provider.UserAgents {
	return &provider.UserAgents{
		BrowserNames:          browserNames,
		WindowsPlatformTokens: windowsPlatformTokens,
		LinuxProcessors:       linuxProcessors,
		MacProcessors:         macProcessors,
		Languages:             languages,
	}
}

var browserNames = []string{
	"firefox", "chrome", "internetExplorer", "opera", "safari", "msedge",
}

var windowsPlatformTokens = []string{
	"Windows NT 6.2", "Windows NT 6.1", "Windows NT 6.0", "Windows NT 5.2", "Windows NT 5.1",
	"Windows NT 5.01", "Windows NT 5.0", "Windows NT 4.0", "Windows 98; Win 9x 4.90", "Windows 98",
	"Windows 95", "Windows CE",
}

var linuxProcessors = []string{
	"i686", "x86_64",
}

var macProcessors = []string{
	"Intel", "PPC", "U; Intel", "U; PPC",
}

var languages = []string{
	"en-US", "sl-SI", "nl-NL",
}
