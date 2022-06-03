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
package gcestartstop

import (
	"context"

	compute "google.golang.org/api/compute/v1"
	//computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func ListInstances(projectId string) (map[string]compute.InstancesScopedList, error) {
	ctx := context.Background()

	computeService, err := compute.NewService(ctx)
	if err != nil {
		return nil, err
	}

	// Create the Google Compute Engine service.
	// instancesClient, err := compute.NewInstancesRESTClient(ctx)
	// if err != nil {
	// 	return fmt.Errorf("NewInstancesRESTClient: %v", err)
	// }
	// defer instancesClient.Close()

	// req := &computepb.AggregatedListInstancesRequest{
	// 	Project:    projectID,
	// 	MaxResults: proto.Uint32(3),
	// }

	// List instances for the project ID.
	//instancesAggregatedListCall := computeService.Instances.AggregatedList(projectId)
	//instancesAggregatedListCall.Filter(strings.Join(filters[:], " "))
	instanceAggregatedList, err := computeService.Instances.AggregatedList(projectId).Do()
	if err != nil {
		return nil, err
	}
	return instanceAggregatedList.Items, nil
}

// func StartInstance(projectId string, instanceName string) (*compute.Operation, error) {
// 	ctx := context.Background()

// 	// Create the Google Compute Engine service.
// 	computeService, err := compute.NewService(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, instance := range ListInstances(projectId) {
// 		if instance.Name == instanceName {
// 			op, err := computeService.Instances.Start(projectId, instanceName, instance.Zone).Do()
// 			if err != nil {
// 				return nil, err
// 			}
// 			return op, nil
// 		}
// 	}
// 	return nil, nil
// }
