package configaccount

// Configuration : Account Configuration File Structure
type Configuration struct {
	Profiles map[string]ProfileConfiguration `json:"profile"`
}

// ProfileConfiguration : Account Profile Configuration
type ProfileConfiguration struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
