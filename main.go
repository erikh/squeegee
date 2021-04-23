package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v3"
)

func errExit(s interface{}) {
	fmt.Fprintln(os.Stderr, s)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		errExit("Please provide a configuration file")
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		errExit(err)
	}

	var config Config

	if err := yaml.Unmarshal(content, &config); err != nil {
		errExit(err)
	}

	for name, gatherer := range config.Metrics {
		spawnURLGatherer(name, gatherer)
	}

	http.Handle("/", promhttp.Handler())
	http.ListenAndServe(":8000", nil)
}
