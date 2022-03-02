// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample listinstances lists the Cloud SQL instances for a given project ID.
package backupinstance

import (
	"context"
	"golang.org/x/oauth2/google"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

func BackupInstance(projectId string, instanceName string) (*sqladmin.Operation, error) {
	ctx := context.Background()

	// Create an http.Client that uses Application Default Credentials.
	hc, err := google.DefaultClient(ctx, sqladmin.SqlserviceAdminScope)
	if err != nil {
		return nil, err
	}

	// Create the Google Cloud SQL service.
	service, err := sqladmin.New(hc)
	if err != nil {
		return nil, err
	}

	//mysetting := &sqladmin.Settings{
	//	ActivationPolicy: "Always",
	//}

	backuprun := &sqladmin.BackupRun{
		Location: "us",
	}
	//if err != nil {
	//	return nil, err
	//}
	op, err := service.BackupRuns.Insert(projectId, instanceName, backuprun).Do()
	//(o *Settings.ActivationPolicy) {
	//	o := "Always"
	//}
	if err != nil {
		return nil, err
	}
	return op, nil
}

