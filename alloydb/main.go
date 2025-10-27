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
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"alloydb.backup.manager/p"
)

func main() {
	// Define command-line flags matching the Parameters struct
	project := flag.String("project", "", "GCP Project ID (Required)")
	location := flag.String("location", "", "GCP Location, e.g., us-central1 (Required)")
	operation := flag.String("operation", "LIST", "Operation to perform: CREATE, DELETE, or LIST (Required)")
	cluster := flag.String("cluster", "ALL", "AlloyDB Cluster ID, or 'ALL' for maintenance operations")
	retention := flag.Int("retention", 365, "Retention period in days (used for DELETE/LIST operations)")

	// Custom usage message
	flag.Usage = func() {
		log.Printf("Usage of %s:\n", os.Args[0])
		log.Println("This tool invokes the AlloyDB backup management function locally.")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Basic validation
	if *project == "" || *location == "" || *operation == "" {
		log.Println("Error: Missing required parameters.")
		flag.Usage()
		os.Exit(1)
	}
	// Construct the parameters object
	params := p.Parameters{
		Project:   *project,
		Location:  *location,
		Operation: *operation,
		Cluster:   *cluster,
		Retention: *retention,
	}

	// Serialize parameters to JSON to mimic the Pub/Sub data payload
	data, err := json.Marshal(params)
	if err != nil {
		log.Fatalf("Failed to marshal parameters to JSON: %v", err)
	}

	// Create the PubSubMessage structure
	msg := p.PubSubMessage{
		Data: data,
	}

	log.Printf("Invoking function with: Operation=%s, Project=%s, Cluster=%s\n", *operation, *project, *cluster)

	// Invoke the function
	ctx := context.Background()
	err = p.ManageAlloyDBBackups(ctx, msg)
	if err != nil {
		log.Fatalf("Function returned error: %v", err)
	}
}
