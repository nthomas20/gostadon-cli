package configapp

// Configuration : Application Configuration File Sructure
// Revision 1
type Configuration struct {
	Revision int                                 `json:"revision"`
	Logging  map[string]LogConfiguration         `json:"logging"`
	Apps     map[string]ApplicationConfiguration `json:"app"`
}

// ApplicationConfiguration : Application Configuration Information
type ApplicationConfiguration struct {
	ServerDomain string                         `json:"server"`
	Name         string                         `json:"name"`
	Type         string                         `json:"type"`
	Scopes       []string                       `json:"scopes"`
	Website      string                         `json:"website"`
	Client       ApplicationClientConfiguration `json:"client"`
}

// ApplicationClientConfiguration : Client Configuration Information
type ApplicationClientConfiguration struct {
	ID          string `json:"id"`
	Secret      string `json:"secret"`
	Token       string `json:"token"`
	RedirectURI string `json:"redirect_uri"`
}

// LogConfiguration : Log Configuration
type LogConfiguration struct {
	Filename string `json:"filename"`
	MaxBytes int64  `json:"max_bytes"`
	MaxFiles int    `json:"max_files"`
}
