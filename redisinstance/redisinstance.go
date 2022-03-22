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
package redisinstance

import (
	"context"
	"log"
	//"golang.org/x/oauth2/google"
	redis "google.golang.org/api/redis/v1"
)

func ListInstances(parent string) (*redis.ListInstancesResponse, error) {
	ctx := context.Background()

	redisService, err := redis.NewService(ctx)

	// Create an http.Client that uses Application Default Credentials.
	// hc, err := google.DefaultClient(ctx, sqladmin.SqlserviceAdminScope)
	// if err != nil {
	// 	return nil, err
	// }

	// // Create the Google Cloud SQL service.
	// service, err := sqladmin.New(hc)
	// if err != nil {
	// 	return nil, err
	// }

	// List instances for the project ID .
	listInstances, err := redisService.Projects.Locations.Instances.List(parent).Do()
	if err != nil {
		log.Println(err)
	}
	return listInstances, nil
}

