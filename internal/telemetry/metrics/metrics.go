package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	system = "nginx-operator"
)

type Metrics struct {
	// Number of NginX deployments per namespace (gauge vector)
	Deployments *prometheus.GaugeVec
	// Number of NginX configmaps per namespace (gauge vector)
	Configmaps *prometheus.GaugeVec
	// Number of failed fetch per namespace (counter vector)
	Failure *prometheus.CounterVec
}

func NewMetrics(namespace string) *Metrics {
	m := &Metrics{}

	m.Deployments = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: system,
	}, []string{"namespace"})
	m.Configmaps = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: system,
	}, []string{"namespace"})
	m.Failure = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: system,
	}, []string{"namespace"})

	if err := prometheus.Register(m.Deployments); err != nil {
		panic(err)
	}

	if err := prometheus.Register(m.Configmaps); err != nil {
		panic(err)
	}

	if err := prometheus.Register(m.Failure); err != nil {
		panic(err)
	}

	return m
}
