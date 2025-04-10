from confluent_kafka import Producer

def delivery_report(err, msg):
    if err is not None:
        print(f"Delivery failed for record {msg.key()}: {err}")
    else:
        print(f"Record {msg.key()} successfully produced to {msg.topic()} [{msg.partition()}] at offset {msg.offset()}")

conf = {
    'bootstrap.servers': 'localhost:9092'
}

producer = Producer(conf)

for x in range(1, 101):
    for y in range(1, 11):
        key = str(x)
        value = str(y)
        producer.produce('test-user', key=key, value=value, callback=delivery_report)
        producer.poll(0)

producer.flush()
