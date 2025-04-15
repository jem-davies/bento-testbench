# rabbitmq

```bash
docker compose up
```

Starts up a rabbitmq management container, the management UI is available on 
http://localhost:15672, username & password is 'guest' and 'guest' navigate to 
"queues and streams", create a queue named 'foo' with the default options. 


You should then be ready to run the bento configs. 

```yaml
input:
  amqp_1:
    urls:
      - amqp://guest:guest@localhost:5672/
    source_address: foo

output:
  amqp_1:
    urls:
      - amqp://guest:guest@localhost:5672/
    target_address: foo
```