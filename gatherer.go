package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/antchfx/jsonquery"
	"github.com/prometheus/client_golang/prometheus"
)

type rt struct {
	http.Transport
	metric MetricSource
}

func (r *rt) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range r.metric.Headers {
		req.Header.Set(key, value)
	}

	return r.Transport.RoundTrip(req)
}

type urlGatherer struct {
	metric MetricSource
	gauge  prometheus.Gauge
}

func spawnURLGatherer(name string, metric MetricSource) {
	fmt.Println("Registering:", name)
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: name,
	})

	prometheus.MustRegister(gauge)

	ug := &urlGatherer{
		metric: metric,
		gauge:  gauge,
	}

	go func(ug *urlGatherer) {
		c := &http.Client{Transport: &rt{metric: metric}}
		ticker := time.NewTicker(ug.metric.Interval)
		defer ticker.Stop()
		for range ticker.C {
			resp, err := c.Get(ug.metric.URL)
			if err != nil {
				fmt.Println(err)
				ug.Set(0)
				continue
			}

			doc, err := jsonquery.Parse(resp.Body)
			if err != nil {
				fmt.Println("couldn't parse", err)
				ug.Set(0)
			} else {
				list := jsonquery.Find(doc, ug.metric.Query)
				if len(list) == 0 {
					fmt.Println("empty list in query:", ug.metric.Query)
					ug.Set(0)
				} else {
					f, err := strconv.ParseFloat(list[0].FirstChild.Data, 64)
					if err != nil {
						ug.Set(0)
					} else {
						ug.Set(f)
					}
				}
			}

			resp.Body.Close()
		}
	}(ug)
}

// Set is a proxy for gauge.Set()
func (ug *urlGatherer) Set(val float64) {
	ug.gauge.Set(val)
}
