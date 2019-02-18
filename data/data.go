package data

import (
	"time"

	types "github.com/chef/code-test/types"
)

var EmptyNodes = types.NodesResponse{
	Nodes: []types.Node{},
	Total: 0,
}

var SampleUnfilteredNodes = types.NodesResponse{
	Nodes: []types.Node{
		types.Node{
			ID:   "e69dc612-7e67-43f2-9b19-256afd385820",
			Name: "my test node",
			Platform: types.Platform{
				Name:    "debian",
				Release: "9.7",
			},
			LastScan: types.LastRun{
				Time:              time.Now(),
				Status:            types.Pass,
				PenultimateStatus: types.Pass,
			},
			LastClientRun: types.LastRun{
				Time:              time.Now(),
				Status:            types.Pass,
				PenultimateStatus: types.Pass,
			},
			Tags: []types.Kv{
				types.Kv{
					Key: "my tag", Value: "is a good one",
				},
			},
		},
		types.Node{
			ID:   "e69dc612-7e67-43f2-9b19-256afd385820",
			Name: "my test node",
			Platform: types.Platform{
				Name:    "debian",
				Release: "9.7",
			},
			LastScan: types.LastRun{
				Time:              time.Now(),
				Status:            types.Pass,
				PenultimateStatus: types.Pass,
			},
			LastClientRun: types.LastRun{
				Time:              time.Now(),
				Status:            types.Pass,
				PenultimateStatus: types.Pass,
			},
			Tags: []types.Kv{
				types.Kv{
					Key: "my tag", Value: "is a good one",
				},
			},
		},
		types.Node{
			ID:   "3d10cc8e-1fae-4544-9cc9-2aa6653cc6b3",
			Name: "another-node",
			Platform: types.Platform{
				Name:    "debian",
				Release: "9.8",
			},
			LastScan: types.LastRun{
				Time:              time.Now(),
				Status:            types.Fail,
				PenultimateStatus: types.Pass,
			},
			LastClientRun: types.LastRun{
				Time:              time.Now(),
				Status:            types.Pass,
				PenultimateStatus: types.Pass,
			},
			Tags: []types.Kv{
				types.Kv{
					Key: "my tag", Value: "is a different one",
				},
			},
		},
		types.Node{
			ID:   "06d167be-8db3-42b2-9274-f94095c3cb93",
			Name: "node",
			Platform: types.Platform{
				Name:    "amazon",
				Release: "2",
			},
			LastScan: types.LastRun{
				Time:              time.Now(),
				Status:            types.Fail,
				PenultimateStatus: types.Pass,
			},
			LastClientRun: types.LastRun{
				Time:              time.Now(),
				Status:            types.Fail,
				PenultimateStatus: types.Pass,
			},
			Tags: []types.Kv{},
		},
		types.Node{
			ID:   "5043c7e7-7e3e-4c7f-9f64-564e3e69e7e2",
			Name: "great-node",
			Platform: types.Platform{
				Name:    "amazon",
				Release: "2",
			},
			LastScan: types.LastRun{
				Time:              time.Now(),
				Status:            types.Pass,
				PenultimateStatus: types.Pass,
			},
			LastClientRun: types.LastRun{
				Time:              time.Now(),
				Status:            types.Fail,
				PenultimateStatus: types.Pass,
			},
			Tags: []types.Kv{
				types.Kv{
					Key: "my tag", Value: "is a good one",
				},
			},
		},
		types.Node{
			ID:   "c8cb0b0d-9df7-4540-ae55-9c9d36684afa",
			Name: "east-node",
			Platform: types.Platform{
				Name:    "debian",
				Release: "9.6",
			},
			LastScan: types.LastRun{
				Time:              time.Now(),
				Status:            types.Pass,
				PenultimateStatus: types.Fail,
			},
			LastClientRun: types.LastRun{
				Time:              time.Now(),
				Status:            types.Pass,
				PenultimateStatus: types.Fail,
			},
			Tags: []types.Kv{},
		},
	},
	Total: 6,
}
