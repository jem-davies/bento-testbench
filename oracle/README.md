# oracle

```bash
docker compose up
```

Starts up a oracle sql container.

Connect via dbeaver: 

host: localhost
port: 1521
database: FREEPDB1 
username: my_user
password: password


### Working bento config: 

```yaml
input:
  generate:
    mapping: root = {"TID":"1001", "AGE":12, "F1":12.66}
    count: 1

output:
  sql_insert:
    driver: oracle
    dsn: oracle://my_user:password@127.0.0.1:1521/FREEPDB1
    table: test
    columns: [ TID, AGE, F1 ]
    args_mapping: |
      root = [ this.TID, this.AGE, this.F1 ]
    max_in_flight: 1
    init_statement: CREATE TABLE test (TID varchar(255), AGE int, F1 float)
    batching:
      count: 500
      period: 100ms
```
