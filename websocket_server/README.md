# Websocket Server

```bash
go run main.go 
```

Opens an endpoint `ws://localhost:8080/ws` that will send a message every 2 seconds - and received messages will be printed to stdout.

### Basic Bento Config

```yaml
input:
  websocket:
    url: ws://localhost:8080/ws

output:
  websocket:
    url: ws://localhost:8080/ws
```