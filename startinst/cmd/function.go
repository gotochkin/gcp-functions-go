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
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
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
func SQLStartStop(ctx context.Context, m PubSubMessage) error {
	var par Parameters 
	err := json.Unmarshal(m.Data,&par) 
	if err != nil {
		log.Println(err)
	}
	//log.Println(string(m.Data))
	log.Println(string(par.Project))
	log.Println(string(par.Operation))
	// Create context
	sqlService, err := sqladmin.NewService(ctx)
	// List instances for the project ID.
	listInstances, err := sqlService.Instances.List(par.Project).Do()
	if err != nil {
		log.Println(err)
	}
	for _, instance := range listInstances.Items {
		fmt.Println(instance.Name)
		if par.Operation == "stop" {
			mysetting := &sqladmin.Settings{
				ActivationPolicy: "Never",
			}
			inst := &sqladmin.DatabaseInstance{
				Settings: mysetting,
			}
			op, err := sqlService.Instances.Patch(par.Project, instance.Name, inst).Do()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(op.Status)
		}
		if par.Operation == "start" {
			mysetting := &sqladmin.Settings{
				ActivationPolicy: "Always",
			}
			inst := &sqladmin.DatabaseInstance{
				Settings: mysetting,
			}
			op, err := sqlService.Instances.Patch(par.Project, instance.Name, inst).Do()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(op.Status)
		}
	}
	fmt.Println("Done")
	return nil
}