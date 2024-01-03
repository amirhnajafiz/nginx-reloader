package metrics

type Metrics struct {
	// Number of NginX deployments per namespace (gauge vector)
	// Number of NginX configmaps per namespace (gauge vector)
	// Number of failed fetch per namespace (counter vector)
	// Operation response time (histogram vector)
}
