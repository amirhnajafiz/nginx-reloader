package config

import "github.com/amirhnajafiz/nginx-configmap-operator/internal/telemetry"

func Default() Config {
	return Config{
		MetricsBindAddress:     8080,
		HealthProbeBindAddress: 8081,
		LeaderElect:            false,
		Telemetry:              telemetry.Config{},
	}
}
