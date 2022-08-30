package config

type SeedUser struct {
	Name        string `yaml:"name"`
	Email       string `yaml:"email"`
	Password    string `yaml:"password"`
	IsSuperuser bool   `yaml:"isSuperuser"`
}

type Seed struct {
	Enabled bool       `yaml:"enabled" conf:"default:false"`
	Users   []SeedUser `yaml:"users"`
}
