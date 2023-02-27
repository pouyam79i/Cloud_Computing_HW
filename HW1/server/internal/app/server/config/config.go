package config

type ServerInfo struct {
	Name  string `json:"name"`
	IP    string `json:"ip"`
	Port  string `json:"port"`
	Debug bool   `json:"debug"`
}
