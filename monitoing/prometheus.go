package monitoing

import (
	"context"
	"fmt"
	"log"
	"time"

	api "github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	model "github.com/prometheus/common/model"
)

// PrometheusQuery runs a prometheus query & returns mode.Value which can be type casted to model.Vector or model.Matrix depending the query
func PrometheusQuery(ctx context.Context, query string) (model.Value, error) {
	client, err := api.NewClient(api.Config{
		Address: "http://127.0.0.1:9090",
	})
	if err != nil {
		log.Printf("Error creating Prometheus API client err: %v", err)
		return nil, err
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	log.Println("Running prometheus query", "query", query)
	result, warnings, err := v1api.Query(ctx, query, time.Now(), v1.WithTimeout(5*time.Second))
	if err != nil {
		log.Println(err, "Error querying Prometheus")
		return nil, err
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings from Prometheus query: %v\n", warnings)
	}
	// log.Println("Prometheus query results", "result", result)
	return result, err
}
