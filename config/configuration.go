package config

type SolarStaffConfig struct {
	URL      string `yaml:"url"`
	ClientID string `yaml:"client_id"`
	Salt     string `yaml:"salt"`
}
