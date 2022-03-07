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
// HelloPubSub consumes a Pub/Sub message.
func ResizeFunc(ctx context.Context, m PubSubMessage) error {
	var par Parameters 
	err := json.Unmarshal(m.Data,&par) 
	if err != nil {
		log.Println(err)
	}
	//log.Println(string(m.Data))
	log.Println(string(par.Project))
	log.Println(string(par.Cluster))
	//ctx := context.Background()
	containerService, err := container.NewService(ctx)
	// sizeRequest := &container.SetNodePoolSizeRequest {
	// 	NodeCount: size,
	// }
	parent := fmt.Sprintf("projects/%s/locations/-", par.Project)
	clusters, err := containerService.Projects.Locations.Clusters.List(parent).Do()
	if err != nil {
		return nil, err
	}
	for _, cluster := range clusters {
		fmt.Println(cluster.Name)
	}
	fmt.Println(parent)
	return nil
}