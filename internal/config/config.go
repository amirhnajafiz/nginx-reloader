package config

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/nginx-configmap-operator/internal/telemetry"

	"github.com/spf13/viper"
)

type Config struct {
	MetricsBindAddress     int              `mapstructure:"metrics_bind_address"`
	HealthProbeBindAddress int              `mapstructure:"health_probe_bind_address"`
	LeaderElect            bool             `mapstructure:"leader_elect"`
	Telemetry              telemetry.Config `mapstructure:"telemetry"`
}

func Load(name string) Config {
	instance := viper.New()
	cfg := Default()

	instance.SetConfigName(name)
	instance.SetConfigType("yaml")
	instance.AddConfigPath(".")

	if err := instance.ReadInConfig(); err != nil {
		log.Println(fmt.Errorf("failed to read config file: %w", err))
	}

	if err := instance.Unmarshal(&cfg); err != nil {
		log.Println(fmt.Errorf("failed to unmarshal into struct: %w", err))
	}

	return cfg
}
