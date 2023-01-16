DB設定更新したら

```sh
cd docker-entrypoint-initdb.d
pg_dump -s -h localhost -U postgres doc_db > test.sql

#または

pg_dump -s -h localhost -U postgres doc_db > docker-entrypoint-initdb.d/test.sql
```