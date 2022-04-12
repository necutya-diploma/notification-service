package config

type Config struct {
	Env   string      `json:"env"`
	GRPC  GRPCConfig  `json:"grpc"`
	Email EmailConfig `json:"email"`
}

type GRPCConfig struct {
	Host string `json:"host" env:"GRPC_HOST"`
	Port int    `json:"port" env:"GRPC_PORT"`
}

type EmailConfig struct {
	Host     string `json:"host" env:"EMAIL_HOST"`
	Port     int    `json:"port" env:"EMAIL_PORT"`
	Username string `json:"username" env:"EMAIL_USERNAME"`
	Password string `json:"password" env:"EMAIL_PASSWORD"`
}

func (c Config) IsDev() bool {
	return c.Env == devEnv
}
