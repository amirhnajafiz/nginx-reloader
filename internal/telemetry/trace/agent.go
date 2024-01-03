package trace

import "context"

const (
	AgentConsoleExporter = "console"
	AgentJaegerExporter  = "jaeger"
)

type Agent struct {
	Type     string  `mapstructure:"type"`
	Endpoint string  `mapstructure:"endpoint"`
	Ratio    float64 `mapstructure:"ratio"`
}

func (a Agent) getExporter(ctx context.Context) {

}
