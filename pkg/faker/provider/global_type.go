package provider

type Global struct {
	Colors       *Colors
	Files        *Files
	Images       *Images
	Internets    *Internets
	Lorems       *Lorems
	Medicals     *Medicals
	Payments     *Payments
	PhoneNumbers *PhoneNumbers
	UserAgents   *UserAgents
	// NOTICE: All fields name should be PLURAL
}

type Colors struct {
	SafeColorNames []string
	AllColorNames  []string
}

type Files struct {
	// REFACTOR: MIMETypes should be `MIMETypesAndExtensions`?
	// MIMETypes type should be map[string][]string
	// because of type restriction, it is set as map[any][]any
	MIMETypes map[any][]any
}

type Images struct{}

type Internets struct {
	// username
	FirstNames      []string
	LastNames       []string
	UserNameFormats []string
	CreateUserName  func(any) any
	// email
	EmailFormats []string
	TLD          []string
	CreateEmail  func(any) any
	//
	URLFormats    []string
	LocalIPBlocks [][]string
}

type Lorems struct {
	Words []string
}

type Medicals struct {
	BloodTypes     []string
	BloodRhFactors []string
}

type Payments struct {
	CardVendors []string
	CardParams  map[string][]string
	IBANFormats map[string][][2]any // each element is [charClass(string), count(int)]
}

// PhoneNumbers holds phone number format data.
// Used by both Global (E164Formats) and Localized (Formats).
type PhoneNumbers struct {
	Formats     []string // locale-specific phone number formats
	E164Formats []string // E.164 international phone number formats
}

type UserAgents struct {
	BrowserNames          []string
	WindowsPlatformTokens []string
	LinuxProcessors       []string
	MacProcessors         []string
	Languages             []string
}
