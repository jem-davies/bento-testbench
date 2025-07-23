import http.client
import json

conn = http.client.HTTPConnection("localhost", 4195)
headers = {
    "Content-Type": "application/json"
}
data = {
    "hello": "world"
}

json_data = json.dumps(data)

conn.request("POST", "/post", body=json_data, headers=headers)

response = conn.getresponse()
print(f"Status code: {response.status}")
print(" ")
print(f"Response Body: {response.read().decode()}")

conn.close()
