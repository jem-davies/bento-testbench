
```sh
export NAME=jem && bento streams
```

```sh
curl http://localhost:4195/streams/foo -X POST --data-binary "@${PWD}/config.yaml"
```

```sh
curl http://localhost:4195/streams/foo -X GET | jq '.config.input.generate.mapping'
```