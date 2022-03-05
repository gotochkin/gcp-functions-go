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
package gkeresize

import (
	"context"
	"golang.org/x/oauth2/google"
	sqladmin "google.golang.org/api/container/v1beta1"
)

func ListClusters(projectId string) ([]*container.Cluster, error) {
	ctx := context.Background()

	// Create the Google Cloud GKE service that uses Application Default Credentials.
	service, err := container.NewService(ctx)
	if err != nil {
		return nil, err
	}

	// List clusters for the project ID.
	clusters, err := service.ProjectsLocationsClustersService.List(projectId).Do()
	if err != nil {
		return nil, err
	}
	return clusters.ListClustersResponse, nil
}