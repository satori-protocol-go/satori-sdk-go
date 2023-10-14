package config

type SatoriApiConfig struct {
	Type        string `json:"type" yaml:"type"`         // http, websocket
	Endpoint    string `json:"endpoint" yaml:"endpoint"` // http://
	Secret      string `json:"secret" yaml:"secret"`
	AccessToken string `json:"access_token" yaml:"access-token"`
	Timeout     int64  `json:"timeout" yaml:"timeout"`         // ms
	PostFormat  string `json:"post_format" yaml:"post-format"` // string array
}

type SatoriEventConfig struct {
	Type        string `json:"type" yaml:"type"` // http-reverse,websocket,websocket-reverse
	Addr        string `json:"addr" yaml:"addr"` // http://\
	Secret      string `json:"secret" yaml:"secret"`
	AccessToken string `json:"access_token" yaml:"access-token"`
	PostFormat  string `json:"post_format" yaml:"post-format"` // string array
}

type SatoriConfig struct {
	Api   SatoriApiConfig   `json:"api" yaml:"api"`
	Event SatoriEventConfig `json:"event" yaml:"event"`
}
