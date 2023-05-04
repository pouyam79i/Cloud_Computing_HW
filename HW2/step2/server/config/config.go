package config

// TODO: remove this hard code add it to env variable!
const (
	RebrandlyURL = "https://api.rebrandly.com/v1/links"
	API_KEY = "f519a4c4afdc4512bcbb8d59bd5024fa"
)

// req for my server
type RequestAPI struct {
	URL string `json:"url"`
}

// res for my server
type ResponseAPI struct {
	LongURL  string `json:"longUrl"`
	ShortURL string `json:"shortUrl"`
	IsCached bool   `json:"isCached"`
	Hostname string `json:"hostname"`
}

// req for rebrandly
type DomainProp struct {
	FullName string `json:"fullName"`
}
type RebrandlyRequestAPI struct {
	Destination string     `json:"destination"`
	Domain      DomainProp `json:"domain"`
}
