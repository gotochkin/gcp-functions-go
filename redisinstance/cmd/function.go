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
	redis "google.golang.org/api/redis/v1"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}
type Parameters struct {
	Project string `json:"project"`
	Instances string `json:"instances"`
}
// RedisFunc consumes a Pub/Sub message.
func RedisFunc(ctx context.Context, m PubSubMessage) error {
	var par Parameters 
	err := json.Unmarshal(m.Data,&par) 
	if err != nil {
		log.Println(err)
	}
	//log.Println(string(m.Data))
	log.Println(string(par.Project))
	log.Println(string(par.Instances))
	// Create context
	redisService, err := redis.NewService(ctx)
	parent := fmt.Sprintf("projects/%s/locations/-", par.Project)
	listInstances, err := redisService.Projects.Locations.Instances.List(parent).Do()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(len(listInstances.Instances))
	if par.Instances == "ALL" {
		for _, instance := range listInstances.Instances {
			parentdelete := fmt.Sprintf("projects/%s/locations/%s/instances/%s", par.Project,instance.LocationId,instance.Name)
			fmt.Println(parentdelete)
			instancedelete, err := redisService.Projects.Locations.Instances.Delete(parentdelete).Do()
			if err != nil {
				log.Println(err)
			}
			fmt.Println(instancedelete)
		}
		fmt.Println("Deleted")
	}
	fmt.Println("Done")
	return nil
}