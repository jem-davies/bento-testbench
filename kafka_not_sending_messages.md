## Setup

### Boot up the containers:

```bash
cd ./postgres
docker compose up -d 
cd ../kafka
docker compose up -d
```

Wait for cluster to go healthy at http://localhost:9021/clusters


### Populate the test-user topic

```bash
cd ./gen_data
uv run main.py
```

### Start the two Bento instances: 

```bash
bento -c config.yaml 

# then after a few seconds in a new terminal:
bento -c config.yaml
```

Then the first bento instance will hang :/