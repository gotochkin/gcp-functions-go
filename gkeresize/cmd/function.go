// Copyright 2022 Gleb Otochkin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"encoding/json"
	"context"
	"fmt"
	"log"
	container "google.golang.org/api/container/v1beta1"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}
type Parameters struct {
	Project string `json:"project"`
	Cluster string `json:"cluster"`
	Size int64 `json: size`
}
// ResizeFunc consumes a Pub/Sub message.
func ResizeFunc(ctx context.Context, m PubSubMessage) error {
	var par Parameters 
	err := json.Unmarshal(m.Data,&par) 
	if err != nil {
		log.Println(err)
	}
	//log.Println(string(m.Data))
	log.Println(string(par.Project))
	log.Println(string(par.Cluster))
	log.Println(string(par.Size))
	// Create context
	containerService, err := container.NewService(ctx)
	sizeRequest := &container.SetNodePoolSizeRequest {
		NodeCount: par.Size,
	}
	parent := fmt.Sprintf("projects/%s/locations/-", par.Project)
	listClusters, err := containerService.Projects.Locations.Clusters.List(parent).Do()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(len(listClusters.Clusters))
	for _, cluster := range listClusters.Clusters {
		if cluster.Name == par.Cluster {
			fmt.Println("cluster name:",cluster.Name,
			            "Node count:", cluster.CurrentNodeCount,
		    )
		    for _, nodepool := range cluster.NodePools {
			    fmt.Println("Node Pool Name:",nodepool.Name)
				parentsize := fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", 
				                          par.Project,cluster.Location,cluster.Name,nodepool.Name)
				nodesize, err := containerService.Projects.Locations.Clusters.NodePools.SetSize(parentsize,sizeRequest).Do()
				if err != nil {
					log.Println(err)
				}
				fmt.Println(nodesize)
		    }
		}		
	}
	fmt.Println("Done")
	return nil
}