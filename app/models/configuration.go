package models

// Configuration : Configuration structure
type Configuration struct {
	Revision       int                                         `json:"revision"`
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
