package config

type SatoriApiConfig struct {
	Version     string `json:"version" yaml:"version"`
	Type        string `json:"type" yaml:"type"`         // http, websocket
	Endpoint    string `json:"endpoint" yaml:"endpoint"` // http://
	Secret      string `json:"secret" yaml:"secret"`
	AccessToken string `json:"access_token" yaml:"access-token"`
	Timeout     int64  `json:"timeout" yaml:"timeout"` // ms
	Platform    string `json:"platform" yaml:"platform"`
	SelfId      string `json:"self_id" yaml:"self_id"`
}

type SatoriEventConfig struct {
	Version     string `json:"version" yaml:"version"`
	Type        string `json:"type" yaml:"type"` // webhook,websocket
	Addr        string `json:"addr" yaml:"addr"` // http://
	Secret      string `json:"secret" yaml:"secret"`
	AccessToken string `json:"access_token" yaml:"access-token"`
}

type SatoriConfig struct {
	Api   SatoriApiConfig   `json:"api" yaml:"api"`
	Event SatoriEventConfig `json:"event" yaml:"event"`
}
