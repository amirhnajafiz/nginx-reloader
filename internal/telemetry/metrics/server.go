package metrics

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Server(port int) {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Println(fmt.Errorf("failed to start promhttp: %w", err))
	}
}
