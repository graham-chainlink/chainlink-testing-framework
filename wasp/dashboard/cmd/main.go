package main

import (
	"github.com/smartcontractkit/chainlink-testing-framework/wasp/dashboard"
)

// main initializes a default dashboard without non-functional requirements (NFRs) or extensions.
// It creates a new dashboard instance and deploys it, using environment variables for configuration.
// If an error occurs during dashboard creation or deployment, the function panics.
func main() {
	// just default dashboard, no NFRs, no dashboard extensions
	// see examples/alerts.go for an example extension
	d, err := dashboard.NewDashboard(nil, nil)
	if err != nil {
		panic(err)
	}
	// set env vars
	//export GRAFANA_URL=...
	//export GRAFANA_TOKEN=...
	//export DATA_SOURCE_NAME=Loki
	//export DASHBOARD_FOLDER=LoadTests
	//export DASHBOARD_NAME=Wasp
	if _, err := d.Deploy(); err != nil {
		panic(err)
	}
}
