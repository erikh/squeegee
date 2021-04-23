package main

import "time"

// Config is the configuration
type Config struct {
	Listen  string                  `yaml:"listen"`
	Metrics map[string]MetricSource `yaml:"metrics"`
}

// MetricSource is the source of the metric
type MetricSource struct {
	URL      string            `yaml:"url"`
	Headers  map[string]string `yaml:"headers"`
	Query    string            `yaml:"query"`
	Interval time.Duration     `yaml:"interval"`
}
