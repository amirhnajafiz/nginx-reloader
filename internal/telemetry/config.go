package telemetry

import (
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/telemetry/metrics"
	"github.com/amirhnajafiz/nginx-configmap-operator/internal/telemetry/trace"
)

type Config struct {
	Metrics metrics.Config
	Tracing trace.Config
}
