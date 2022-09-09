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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/oauth2/google"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}
type Parameters struct {
	Project   string `json:"project"`
	Location  string `json:"location"`
	Operation string `json:"operation"`
	Cluster   string `json:"cluster"`
	Retention int    `json:"retention"`
}

type Backups struct {
	Backups []Backup `json:"backups"`
}

type Backup struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Uid         string `json:"uid"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	State       string `json:"state"`
	DeleteTime  string `json:"deleteTime"`
	Description string `json:"description"`
	ClusterName string `json:"clusterName"`
	Reconciling bool   `json:"reconciling"`
	Etag        string `json:"etag"`
}

func DeleteAlloyDBBackups(ctx context.Context, m PubSubMessage) error {
	//Parameters
	var par Parameters
	err := json.Unmarshal(m.Data, &par)
	if err != nil {
		log.Println(err)
	}
	project := par.Project
	location := par.Location
	state := "READY"
	cluster := par.Cluster
	retention := par.Retention
	retention = -(retention)
	backupsUrl := "https://alloydb.googleapis.com/v1beta/projects/" + project + "/locations/" + location + "/backups"
	// Get default credentials
	//ctx := context.Background()
	client, err := google.DefaultClient(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		log.Fatal(err)
	}
	listBackups, err := http.NewRequest("GET", backupsUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(listBackups)
	if err != nil {
		log.Fatal(err)
	}
	listBackupBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//
	backups := Backups{}
	err = json.Unmarshal(listBackupBody, &backups)
	t1 := time.Now().AddDate(0, 0, retention)
	for _, backup := range backups.Backups {
		t2, err := time.Parse(time.RFC3339, backup.CreateTime)
		if err != nil {
			log.Fatal(err)
		}
		if t1.Sub(t2) > 0 && backup.State == state {
			if cluster == "ALL" {
				// Debug
				// fmt.Println(backup.ClusterName)
				// fmt.Println(backup.CreateTime)
				// fmt.Println(t2)
				// fmt.Println(t1)
				// fmt.Println(backup.Name)
				// Delete the backup
				deleteUrl := "https://alloydb.googleapis.com/v1beta/" + backup.Name
				deleteBackup, err := http.NewRequest("DELETE", deleteUrl, nil)
				if err != nil {
					log.Fatal(err)
				}
				resp, err := client.Do(deleteBackup)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(resp)

			} else if cluster == strings.Split(backup.ClusterName, "/")[5] {
				// Debug
				// fmt.Println(backup.ClusterName)
				// fmt.Println(backup.CreateTime)
				// fmt.Println(t2)
				// fmt.Println(t1)
				// fmt.Println(backup.Name)
				// Delete the backup
				deleteUrl := "https://alloydb.googleapis.com/v1beta/" + backup.Name
				deleteBackup, err := http.NewRequest(par.Operation, deleteUrl, nil)
				if err != nil {
					log.Fatal(err)
				}
				resp, err := client.Do(deleteBackup)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(resp)
			}

		}
	}
	fmt.Println("Done")
	return nil
}
