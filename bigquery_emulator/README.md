## 

```bash
docker run -p 9050:9050 -p 9060:9060 ghcr.io/goccy/bigquery-emulator --project=test-project --dataset=test-dataset
```

```bash
curl -X POST http://localhost:9050/projects/test-project/datasets/test-dataset/tables \
-H "Content-Type: application/json" \
-d '{
  "tableReference": {
    "projectId": "test-project",
    "datasetId": "test-dataset",
    "tableId": "test_table"
  },
  "schema": {
    "fields": [
      { "name": "first_name", "type": "STRING" },
      { "name": "last_name", "type": "STRING" },
      { "name": "age", "type": "INTEGER" }
    ]
  }
}'
```

### Bento Config:

```yaml
output:
  gcp_bigquery_write_api:
    project: test-project
    dataset: test-dataset
    table: test_table
    # message_format: protobuf
    endpoint: 
      http: http://localhost:9050
      grpc: 0.0.0.0:9060
```

### Query Table

```bash
curl -X POST http://localhost:9050/projects/test-project/queries \
-H "Content-Type: application/json" \
-d '{
  "query": "SELECT * FROM `test-project.test-dataset.test_table`"
}'
```