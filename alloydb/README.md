# AlloyDB Backup Management CLI

This tool allows you to invoke the AlloyDB backup management Google Cloud Function locally from the command line for testing and operational purposes.

## Setup

1.  Ensure Go is installed.
2.  Initialize the module and dependencies:
    ```bash
    go mod init alloydb.backup.manager
    go mod tidy
    ```
3.  Authenticate with Google Cloud:
    ```bash
    gcloud auth application-default login
    ```
    *Alternatively, set the `GOOGLE_APPLICATION_CREDENTIALS` environment variable to the path of a service account key file.*

## Usage

Run the tool using `go run main.go` followed by the required flags.

### Available Flags

| Flag | Description | Required? |
| :--- | :--- | :--- |
| `-project` | Google Cloud Project ID. | Yes |
| `-location`| GCP Region (e.g., `us-central1`). | Yes |
| `-operation`| Action to perform: `CREATE`, `DELETE`, `LIST`. | Yes |
| `-cluster` | AlloyDB Cluster ID. Use `ALL` for maintenance ops across all clusters in the location. | Yes (for CREATE) |
| `-retention`| Retention days for cleaning up old backups. Defaults to 30. | No (optional) |

### Examples

#### 1. Create an On-Demand Backup
Creates an immediate backup for a specific cluster.

```bash
go run main.go \
  -project=my-gcp-project-id \
  -location=us-central1 \
  -cluster=my-alloydb-cluster \
  -operation=CREATE
```

#### 2. Delete Old Backups (Retention policy)
Delete backups older than 60 days in the specific project and location
```
go run main.go \
  -project=my-gcp-project-id \
  -location=us-central1 \
  -cluster=ALL \
  -retention=60 \
  -operation=DELETE
```

#### 3. Delete Old backupds for a Specific Cluster
Delete backups older than 60 days for a cluster 
```
go run main.go \
  -project=my-gcp-project-id \
  -location=us-central1 \
  -cluster=my-alloydb-cluster \
  -operation=DELETE
```
