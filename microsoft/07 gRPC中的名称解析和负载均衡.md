## grpc客户端如何拿到所有服务器ip地址 name resolving
```
gRPC中默认使用的名称解析器是 DNS，即在gRPC客户端执行grpc.Dial时提供域名，默认会将DNS解析出对应的IP列表返回。
使用默认DNS解析器的名称语法为：dns:[//authority/]host[:port]
```
```
conn, err := grpc.Dial("dns:///localhost:8972",
	grpc.WithTransportCredentials(insecure.NewCredentials()),
)
```
```
grpc客户端通过grpc.WithDefaultServiceConfig来配置要使用的负载均衡策略
```
```
conn, err := grpc.Dial(
	"q1mi:///resolver.liwenzhou.com",
	grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), // 这里设置初始策略
	grpc.WithTransportCredentials(insecure.NewCredentials()),
)
```
```
root@jasper:/home/jasper# docker run -d \
  --name mysql-test \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=123456 \
  mysql:8.0
```
```
docker exec -it mysql-test bash
```
```
bash-5.1# mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
```