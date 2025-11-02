```sh
# boot the server
go run ./test_server/main.go
```

```sh
# start bento 
bento -c ./config.yaml
```

```sh
# send request to bento which will act like a proxy server 
curl -X POST http://localhost:4195/post \
-H "Content-Type: application/json" \
-d '{"name":"Jem"}'
```