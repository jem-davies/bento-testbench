# postgres

```bash
docker compose up
```

Starts up a postgres container and a bento container that will create a `test_data` table:

```sql
CREATE TABLE IF NOT EXISTS test_data (
          id SERIAL PRIMARY KEY,
          first_name VARCHAR(255),
          last_name VARCHAR(255),
          age INT
        );
```

and populate it with a 1000 rows of fake data.

After running `docker compose down` - if you don't want to persist this data then remove the volume at ./postgres/postgres_data, otherwise an additional 1000 rows will be added to `test_data` each time you start the docker compose stack. 

### Basic Bento Config

```yaml
input:
  sql_select:
    driver: postgres
    dsn: postgresql://user:password@localhost:5432/db?sslmode=disable
    table: test_data
    columns:
      - id
      - first_name
      - last_name
      - age

output:
  sql_insert:
    driver: postgres
    dsn: postgresql://user:password@localhost:5432/db?sslmode=disable
    table: test_data
    columns:
      - first_name
      - last_name
      - age
    args_mapping: |
        root = [
          this.first_name,
          this.last_name,
          this.age
        ]
```

