package config

const (
	API_CODEX_URL = "https://api.codex.jaagrav.in"
)

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

type ClientMSG struct {
	Type string `json:"type"`
	Info string `json:"info"`
}

type CodexAPI struct {
	Code     string `json:"code"`
	Language string `json:"language"`
	Input    string `json:"input"`
}

type ClientCode struct {
	Token string   `json:"token"`
	CodeX CodexAPI `json:"codex"`
}

type ResCodeX struct {
	TimeStamp int    `json:"timeStamp"`
	Status    int    `json:"status"`
	Output    string `json:"output"`
	Error     string `json:"error"`
	Language  string `json:"language"`
	Info      string `json:"info"`
}
