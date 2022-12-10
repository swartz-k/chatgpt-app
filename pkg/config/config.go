package config

type Config struct {
	Mode    string `json:"mode"`
	Addr    string `json:"addr"`
	Session string `json:"session"`
}

func newDefaultConfig() *Config {
	return &Config{Addr: "127.0.0.1:8092"}
}
