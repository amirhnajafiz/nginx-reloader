package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	// Number of NginX deployments per namespace (gauge vector)
	Deployments *prometheus.GaugeVec
	// Number of NginX configmaps per namespace (gauge vector)
	Configmaps *prometheus.GaugeVec
	// Number of failed fetch per namespace (counter vector)
	Failure *prometheus.CounterVec
}
