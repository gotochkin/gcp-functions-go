// Copyright 2022 Gleb Otochkin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// / Package p contains a Pub/Sub Cloud Function.
package p

import (
	"bytes"
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
	Debug     bool   `json:"debug"`
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
	ExpiryTime  string `json:"expiryTime"`
	Etag        string `json:"etag"`
}

func operBackup(ctx context.Context, project string, location string, cluster string, operation string, retention int, state string, debug bool) {
	// Get default credentials
	//ctx := context.Background()
	client, err := google.DefaultClient(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		log.Fatal(err)
	}
	backupsUrl := "https://alloydb.googleapis.com/v1beta/projects/" + project + "/locations/" + location + "/backups"
	listBackups, err := http.NewRequest("GET", backupsUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(listBackups)
	if err != nil {
		log.Fatal(err)
	}
	//The io.ReadAll operation can be optimized using different approach
	listBackupBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//
	backups := Backups{}
	if err := json.Unmarshal(listBackupBody, &backups); err != nil {
		log.Fatal(err)
	}
	if len(backups.Backups) == 0 && debug {
		log.Println("Debug: No backups found.")
		log.Printf("Debug: To list backups manually, run the following command:")
		log.Printf("curl -X GET -H \"Authorization: Bearer $(gcloud auth print-access-token)\" %s", backupsUrl)
	}
	t1 := time.Now().AddDate(0, 0, retention)
	fmt.Printf("Searching for backups older than: %v\n", t1.Format("2006-01-02"))
	count := 0
	clusterBackupCount := 0
	for _, backup := range backups.Backups {
		currentClusterId := strings.Split(backup.ClusterName, "/")[5]
		if cluster != "ALL" && cluster != currentClusterId {
			continue
		}
		clusterBackupCount++
		t2, err := time.Parse(time.RFC3339, backup.CreateTime)
		if err != nil {
			log.Printf("Error parsing time for backup %s: %v", backup.Name, err)
			continue
		}
		t3, err := time.Parse(time.RFC3339, backup.ExpiryTime)
		if err != nil {
			log.Printf("Error parsing expiryTime for backup %s: %v", backup.Name, err)
			continue
		}
		if t1.Sub(t2) > 0 && backup.State == state {
			count++
			if operation == "LIST" {
				// Just print details, don't make an API call
				fmt.Printf("[LIST] Found Backup: %s \n\tCluster: %s \n\tCreated: %s\n",
					backup.Name, currentClusterId, backup.CreateTime)
			} else {
				// Handle DELETE (or other operational verbs)
				fmt.Printf("[%s] Performing operation on: %s\n", operation, backup.Name)
				if t1.Sub(t3) < 0 && operation == "DELETE" {
					log.Printf("Cannot %s backup %s: before its expiration day: %v", operation, backup.Name, t3)
					continue
				}
				backupUrl := "https://alloydb.googleapis.com/v1beta/" + backup.Name
				operateBackupReq, err := http.NewRequest(operation, backupUrl, nil)
				if err != nil {
					log.Fatal(err)
				}
				opResp, err := client.Do(operateBackupReq)
				if err != nil {
					log.Printf("Failed to %s backup %s: %v", operation, backup.Name, err)
					continue
				} else if opResp.Status != "200" {
					log.Printf("Failed to %s backup %s: %v", operation, backup.Name, err)
					continue
				}
				// It's good practice to close these response bodies too, even if we don't read them fully
				opResp.Body.Close()
				fmt.Printf("Operation %s response status: %s\n", operation, opResp.Status)
			}
		}
	}
	if debug {
		if clusterBackupCount == 0 && cluster != "ALL" {
			log.Printf("Debug: No backups found for cluster '%s'.", cluster)
		} else if count == 0 && clusterBackupCount > 0 {
			log.Println("Debug: Backups for cluster found, but none are older than the specified retention period.")
			for _, backup := range backups.Backups {
				log.Printf("Debug: Found backup: %s, Created: %s, State: %s", backup.Name, backup.CreateTime, backup.State)
			}
		}
	}
}
func createBackup(ctx context.Context, project string, location string, cluster string, operation string) {
	// Get default credentials
	//ctx := context.Background()
	client, err := google.DefaultClient(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		log.Fatal(err)
	}
	backupsUrl := "https://alloydb.googleapis.com/v1beta/projects/" + project + "/locations/" + location + "/backups/?backupId=on-demand-bkp-" + time.Now().Format("2006-01-02-150405")
	jsonVal := map[string]string{"clusterName": "projects/" + project + "/locations/" + location + "/clusters/" + cluster, "type": "ON_DEMAND"}
	backupJSON, _ := json.Marshal(jsonVal)
	createBackup, err := http.NewRequest("POST", backupsUrl, bytes.NewBuffer(backupJSON))
	createBackup.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(createBackup)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func ManageAlloyDBBackups(ctx context.Context, m PubSubMessage) error {
	//Parameters
	var par Parameters
	err := json.Unmarshal(m.Data, &par)
	if err != nil {
		log.Println(err)
	}
	if par.Project == "" && par.Debug {
		log.Println("Debug: No project specified.")
	}
	project := par.Project
	location := par.Location
	state := "READY"
	cluster := par.Cluster
	operation := par.Operation
	retention := par.Retention
	retention = -(retention)
	if operation == "DELETE" || operation == "LIST" {
		//
		operBackup(ctx, project, location, cluster, operation, retention, state, par.Debug)
	} else if operation == "CREATE" {
		createBackup(ctx, project, location, cluster, operation)
	}

	fmt.Println("Done")
	return nil
}
