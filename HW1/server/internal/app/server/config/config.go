package config

type ServerInfo struct {
	Name  string `json:"name"`
	IP    string `json:"ip"`
	Port  string `json:"port"`
	Debug bool   `json:"debug"`
}

type SingInInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorData struct {
	Type string `json:"type"`
	Info string `json:"info"`
}

type SuccessMSG struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}
