### GCP Spanner


```bash
docker run -p 9010:9010 -p 9020:9020 gcr.io/cloud-spanner-emulator/emulator
```

```bash
go run . # set's up a database & table - and creates data changes 
```

Config for the gcp_spanner_cdc TODO: create a config for gcp_spanner with the sql components?

```bash
export SPANNER_EMULATOR_HOST=localhost:9010
```

Then run: 

```yaml
input:
  gcp_spanner_cdc: 
    spanner_dsn: projects/project-id/instances/instance-id/databases/stream-database
    spanner_metadata_dsn: projects/project-id/instances/instance-id/databases/stream-database
    metadata_table: meta_table
    stream_name: Stream

output:
  file:
    path: ./data.txt
    codec: lines
```