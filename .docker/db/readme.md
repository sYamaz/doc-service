DB設定更新したら

```sh
# コンテナ内

cd docker-entrypoint-initdb.d
pg_dump -s -h localhost -U postgres doc-db > 02_ddl.sql

#または

pg_dump -s -h localhost -U postgres doc-db > docker-entrypoint-initdb.d/02_ddl.sql

# コンテナ外から
docker exec -i doc-api_db ash -c "pg_dump -s -h localhost -U postgres doc-db > docker-entrypoint-initdb.d/02_ddl.sql"
```