# dynamodb

Start the dynamodb-local container & check connection: 

```sh
docker compose up -d
aws dynamodb describe-limits --endpoint-url http://localhost:8000 --region us-east-1
```

create table:

```sh
aws dynamodb create-table \
    --table-name Music \
    --attribute-definitions \
        AttributeName=Artist,AttributeType=S \
        AttributeName=SongTitle,AttributeType=S \
    --key-schema AttributeName=Artist,KeyType=HASH AttributeName=SongTitle,KeyType=RANGE \
    --billing-mode PAY_PER_REQUEST \
    --table-class STANDARD \
    --endpoint-url http://localhost:8000 \
    --region us-east-1 \
```

run bento: 

```sh
bento -c config.yaml
```

reference: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/SettingUp.html

### TTL 

create table: 

```sh
aws dynamodb create-table \
    --table-name Music \
    --attribute-definitions \
        AttributeName=Artist,AttributeType=S \
        AttributeName=SongTitle,AttributeType=S \
    --key-schema AttributeName=Artist,KeyType=HASH AttributeName=SongTitle,KeyType=RANGE \
    --billing-mode PAY_PER_REQUEST \
    --table-class STANDARD \
    --endpoint-url http://localhost:8000 \
    --region us-east-1 
```

enable ttl: 

```sh
aws dynamodb update-time-to-live \
    --table-name Music \
    --time-to-live-specification "Enabled=true, AttributeName=ttl"
    --endpoint-url http://localhost:8000 \
    --region us-east-1 \
```

insert a record using bento: 

```sh
bento -c config_ttl.yaml
```

query the table: 

```sh
aws dynamodb scan \
    --table-name Music \
    --endpoint-url http://localhost:8000 \
    --region us-east-1
```