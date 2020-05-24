package models

// MainConfiguration : Application Configuration File Sructure
// Revision 1
type MainConfiguration struct {
	Revision int                                 `json:"revision"`
	Logging  map[string]LogConfiguration         `json:"logging"`
	Client   map[string]ApplicationConfiguration `json:"client"`
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

// AccountConfiguration : Account Configuration File Structure
type AccountConfiguration struct {
	Profiles map[string]ProfileConfiguration `json:"profile"`
}

// ProfileConfiguration : Account Profile Configuration
type ProfileConfiguration struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
