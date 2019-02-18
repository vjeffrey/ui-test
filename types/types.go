package types

import "time"

// NodesResponse is used for the `/nodes` endpoint response
type NodesResponse struct {
	Nodes []Node
	Total int
}

type Node struct {
	ID            string
	Name          string
	Platform      Platform
	LastScan      LastRun
	LastClientRun LastRun
	Tags          []Kv
}

type Kv struct {
	Key   string
	Value string
}

type Platform struct {
	Name    string
	Release string
}
type LastRun struct {
	Time              time.Time
	Status            string
	PenultimateStatus string
}

type NodesRequest struct {
	Filters []Kv
}

const Pass = "passed"
const Fail = "failed"
const Skip = "skipped"
