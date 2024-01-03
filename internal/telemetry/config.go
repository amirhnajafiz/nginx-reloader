package telemetry

import (
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/telemetry/logger"
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/telemetry/metrics"
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/telemetry/trace"
)

type Config struct {
	Metrics metrics.Config `mapstructure:"metrics"`
	Tracing trace.Config   `mapstructure:"tracing"`
	Logger  logger.Config  `mapstructure:"logger"`
}
