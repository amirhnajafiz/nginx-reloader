package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

// Config struct is used to store reloader parameters.
type Config struct {
	Type         string `env:"NR_TYPE" envDefault:"download"`
	Address      string `env:"NR_ADDRESS,required"`
	NginxHTMLDir string `env:"NR_NGINX_HTML_DIR" envDefault:"/usr/share/nginx/html"`
	TmpLocalDir  string `env:"NR_TMP_LOCAL_DIR" envDefault:"/etc/nginx-reloader/tmp"`
}

// Load env variables into a config struct.
func Load() (*Config, error) {
	instance := &Config{}

	if err := env.Parse(instance); err != nil {
		return nil, fmt.Errorf("failed to parse config values: %v", err)
	}

	return instance, nil
}
