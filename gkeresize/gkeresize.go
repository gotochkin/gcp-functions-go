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
package gkeresize

import (
	"context"
	"golang.org/x/oauth2/google"
	container "google.golang.org/api/container/v1beta1"
)

func ListClusters(parent string) (*container.ListClustersResponse, error) {
	ctx := context.Background()

	// Create an http.Client that uses Application Default Credentials.
	hc, err := google.DefaultClient(ctx, container.CloudPlatformScope)
	if err != nil {
		return nil, err
	}

	// Create the Google Cloud GKE service that uses Application Default Credentials.
	service, err := container.New(hc)
	if err != nil {
		return nil, err
	}
	// Service -> ProjectsService -> ProjectsLocationsService -> List(projectID)  ProjectsLocationsClustersService?
	//cs. err := 

	// List clusters for the project ID.
	
	clusters, err := service.Projects.Locations.Clusters.List(parent).Do()
	if err != nil {
		return nil, err
	}
	return clusters, nil
}

func ResizeClusters(parent string, size int64) (*container.Operation, error) {
	ctx := context.Background()

	// Create an http.Client that uses Application Default Credentials.
	hc, err := google.DefaultClient(ctx, container.CloudPlatformScope)
	if err != nil {
		return nil, err
	}

	// Create the Google Cloud GKE service that uses Application Default Credentials.
	service, err := container.New(hc)
	if err != nil {
		return nil, err
	}
	// Service -> ProjectsService -> ProjectsLocationsService -> List(projectID)  ProjectsLocationsClustersService?
	//cs. err := 

	// List clusters for the project ID.
	sizeRequest := &container.SetNodePoolSizeRequest {
		NodeCount: size,
	}
	
	nodesize, err := service.Projects.Locations.Clusters.NodePools.SetSize(parent,sizeRequest).Do()
	if err != nil {
		return nil, err
	}
	return nodesize, nil
}
