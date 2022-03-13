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
  "gleb.ca/listinstances"
  "fmt"
  "strconv"
  "log"
)
func main() {
  projects := [...]string{
    "gleb-sandbox",
  }
  for _, project := range projects {
    listInstances, err := listinstances.ListInstances(project) 
    if err != nil {
      log.Fatalln(err)
    }
    for i, instance := range listInstances {
      //fmt.Printf("Number: %s instance name: %s\n",strconv.Itoa(i),instance.Name)
      fmt.Println("Number:",strconv.Itoa(i),
                  "instance name:",instance.Name,
                  "State:", instance.State,
                  "Settings.ActivationPolicy:", instance.Settings.ActivationPolicy,
                  "Settings.SettingsVersion:", instance.Settings.SettingsVersion,
                )
      fmt.Printf("%#v\n",instance.BackupConfiguration)
    }
  }
}
