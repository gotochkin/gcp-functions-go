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
	Instance string `json:"instance"`
	Location string `json:"location"`
	Description string `json:"description"`
}
// SQLBackup consumes a Pub/Sub message.
func SQLBackup(ctx context.Context, m PubSubMessage) error {
	var par Parameters 
	err := json.Unmarshal(m.Data,&par) 
	if err != nil {
		log.Println(err)
	}
	//log.Println(string(m.Data))
	log.Println(string(par.Project))
	log.Println(string(par.Instance))
	log.Println(string(par.Location))
	log.Println(string(par.Description))
	// Create context
	sqlService, err := sqladmin.NewService(ctx)
    // 
	backuprun := &sqladmin.BackupRun{
		Location: par.Location,
		Description: par.Description,
	}

	op, err := sqlService.BackupRuns.Insert(par.Project, par.Instance, backuprun).Do()
	if err != nil {
		log.Println(err)
	}
	//
	fmt.Println(op.Status)
	fmt.Println("Done")
	return nil
}