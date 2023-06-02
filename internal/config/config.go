package config

type Config struct {
	sekai  Sekai // Here we will place out config
	interx Interx
}

type Sekai struct {
	config ConfToml
}

type Interx struct {
}

type ConfToml struct {
}

type DockerConfig struct {
	Host       string `json:"Host"`
	APIVersion string `json:"APIVersion"`
	CertPath   string `json:"CertPath"`
	CacertPath string `json:"CacertPath"`
	KeyPath    string `json:"KeyPath"`
}
