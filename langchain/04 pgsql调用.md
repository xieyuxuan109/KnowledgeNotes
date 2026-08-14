```
docker pull postgres:16.4 
docker run -d \
  --name aisec-postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=root \
  -e POSTGRES_DB=aisec \
  -p 5432:5432 \
  -v aisec_pgdata:/var/lib/postgresql/data \
  postgres:16.4
```
## 进入容器
```
docker exec -it aisec-postgres bash
psql -U postgres -d aisec
列举表，相当于mysql的show tables
\dt
```
## 扩展向量
```
apt update
apt install postgresql-15-pgvector
systemctl restart postgresql
psql -U postgres -d deepwiki
CREATE EXTENSION vector;
```