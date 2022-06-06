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
package main

import (
	"fmt"
	"log"
	"strings"

	"gleb.ca/gcestartstop"
	//"strconv"
)

func main() {
	projects := [...]string{
		"gleb-sandbox",
	}
	// instances := [...]string{
	// 	"gleb-sandbox-1",
	// }

	for _, project := range projects {
		// for _, instance := range instances {
		// 	Instance, err := gcestartstop.StartInstance(project, instance)
		// 	if err != nil {
		// 		log.Fatalln(err)
		// 	}
		// 	fmt.Println(Instance)
		// }
		instancesScopedList, err := gcestartstop.ListInstances(project)
		if err != nil {
			log.Fatalln(err)
		}
		//fmt.Println(instancesScopedList)
		for zone, instances := range instancesScopedList {
			if len(instances.Instances) > 0 {
				//fmt.Println(zone)
				//fmt.Println(len(instances.Instances))
				for _, instance := range instances.Instances {
					fmt.Println(instance.Name + " "+instance.Status)
					if instance.Status == "RUNNING" {
						zonename := strings.Split(zone,"/")
						fmt.Println(project + " " + zonename[1] + " " + instance.Name)
						op, err := gcestartstop.StopInstance(project,zonename[1],instance.Name)
						if err != nil {
							log.Fatalln(err)
						}
						fmt.Println(op.Status)
					}
				}
			}
		}
	}
}
