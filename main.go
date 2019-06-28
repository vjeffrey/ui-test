package main

import (
	"context"
	"encoding/json"
	"net/http"

	"code-test/data"
	"code-test/types"

	log "github.com/sirupsen/logrus"
)

func main() {
	serveAddress := "0.0.0.0:2133"

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	r := http.NewServeMux()

	log.Infof("serving /nodes endpoint at %s", serveAddress)
	log.Infof("available filters are %s", "name platform_name platform_release last_scan_status last_client_run_status")
	r.HandleFunc("/nodes", nodesHandler)

	err := http.ListenAndServe(serveAddress, r)
	if err != nil {
		log.Fatal("unable to serve http")
	}
}

func nodesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	decoder := json.NewDecoder(r.Body)
	var t types.NodesRequest
	err := decoder.Decode(&t)
	data := handleFilters(t)
	byteData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("unable to parse data")
	}
	w.Write(byteData)
}

func handleFilters(request types.NodesRequest) types.NodesResponse {
	log.Infof("received filters request: %v", request.Filters)
	nodesData := make([]types.Node, 0)
	unfilteredNodesData := data.SampleUnfilteredNodes
	if len(request.Filters) == 0 {
		log.Infof("no filters included in query. sending back all nodes")
		return unfilteredNodesData
	}
	for _, filter := range request.Filters {
		nodesData = append(nodesData, returnMatchingItems(unfilteredNodesData.Nodes, filter)...)
	}
	return types.NodesResponse{Nodes: nodesData, Total: len(nodesData)}
}

func returnMatchingItems(nodesArr []types.Node, filter types.Kv) []types.Node {
	matchingNodes := make([]types.Node, 0)
	for _, node := range nodesArr {

		switch filter.Key {
		case "name":
			if node.Name == filter.Value {
				matchingNodes = append(matchingNodes, node)
			}
		case "platform_name":
			if node.Platform.Name == filter.Value {
				matchingNodes = append(matchingNodes, node)
			}
		case "platform_release":
			if node.Platform.Release == filter.Value {
				matchingNodes = append(matchingNodes, node)
			}
		case "last_scan_status":
			if node.LastScan.Status == filter.Value {
				matchingNodes = append(matchingNodes, node)
			}
		case "last_client_run_status":
			if node.LastClientRun.Status == filter.Value {
				matchingNodes = append(matchingNodes, node)
			}
		default:
			log.Error("invalid filter field.")
		}
	}
	return matchingNodes
}
