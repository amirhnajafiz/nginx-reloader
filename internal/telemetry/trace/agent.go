package trace

import "context"

const (
	AgentConsoleExporter = "console"
	AgentJaegerExporter  = "jaeger"
)

type Agent struct {
	Type     string
	Endpoint string
	Ratio    float64
}

func (a Agent) getExporter(ctx context.Context) {

}
