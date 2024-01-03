package config

import "github.com/amirhnajafiz/nginx-configmap-operator/internal/telemetry"

type Config struct {
	MetricsBindAddress     int              `mapstructure:"metrics_bind_address"`
	HealthProbeBindAddress int              `mapstructure:"health_probe_bind_address"`
	Telemetry              telemetry.Config `mapstructure:"telemetry"`
}
