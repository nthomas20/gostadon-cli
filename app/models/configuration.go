package models

// Configuration : Application Configuration File Sructure
// Revision 1
type Configuration struct {
	Revision       int                                         `json:"revision"`
	Logging        map[string]LogConfiguration                 `json:"logging"`
	MastodonClient map[string]MastodonApplicationConfiguration `json:"mastodon_client"`
}

// MastodonApplicationConfiguration : Application Configuration Information
type MastodonApplicationConfiguration struct {
	ServerDomain string                      `json:"server"`
	Name         string                      `json:"app_name"`
	Scopes       []string                    `json:"scopes"`
	Website      string                      `json:"website"`
	Client       MastadonClientConfiguration `json:"client"`
}

// MastadonClientConfiguration : Client Configuration Information
type MastadonClientConfiguration struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
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
