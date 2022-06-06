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
	"strings"
	"fmt"
	"log"
	compute "google.golang.org/api/compute/v1"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}
type Parameters struct {
	Project string `json:"project"`
	Operation string `json:"operation"`
}
// SQLStartStop consumes a Pub/Sub message.
func GCEStartStop(ctx context.Context, m PubSubMessage) error {
	var par Parameters 
	err := json.Unmarshal(m.Data,&par) 
	if err != nil {
		log.Println(err)
	}
	//log.Println(string(m.Data))
	log.Println(string(par.Project))
	log.Println(string(par.Operation))
	// Create context
	computeService, err := compute.NewService(ctx)
	// Aggregated list instances for the project ID.
	instanceAggregatedList, err := computeService.Instances.AggregatedList(par.Project).Do()
	if err != nil {
		log.Println(err)
	}
	for zone, instances := range instanceAggregatedList.Items {
		//
		if len(instances.Instances) > 0 {
			zonename := strings.Split(zone,"/")
			for _, instance := range instances.Instances {
				//fmt.Println(instance.Name + " " + instance.Status)
				if instance.Status == "RUNNING" && par.Operation == "Stop" {
					//fmt.Println(project + " " + zonename[1] + " " + instance.Name)
					op, err := computeService.Instances.Stop(par.Project,zonename[1],instance.Name).Do()
					if err !=nil {
						log.Println(err)
					}
					fmt.Println(op.Status)
				} else if instance.Status == "TERMINATED" && par.Operation == "Start" {
					//fmt.Println(project + " " + zonename[1] + " " + instance.Name)
					op, err := computeService.Instances.Start(par.Project,zonename[1],instance.Name).Do()
					if err !=nil {
						log.Println(err)
					}
					fmt.Println(op.Status)
				}
			}
		}
	}
	
	fmt.Println("Done")
	return nil
}