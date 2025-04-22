# Nats

### output_kv

```bash
brew tap nats-io/nats-tools
brew install nats-io/nats-tools/nats
docker compose up 
nats kv add my-kv
bento -c output_kv.yaml
nats kv get my-kv foo
```