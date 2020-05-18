package cmd

// ApplicationConfiguration : Application Configuration Information
type ApplicationConfiguration struct {
	ServerDomain string              `json:"domain"`
	Name         string              `json:"app_name"`
	Scopes       []string            `json:"scopes"`
	Website      string              `json:"website"`
	Client       ClientConfiguration `json:"client"`
}

// ClientConfiguration : Client Configuration Information
type ClientConfiguration struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
}
