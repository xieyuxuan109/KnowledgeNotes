```
配置
vim /etc/hosts
10.17.210.200 nexus.jasper.cn.com:8082
```
```
docker配置
vim /etc/docker/daemon.json
```
```
{
  "insecure-registries": ["10.17.210.200:8082"]
}
```
```
重启docker
systemctl daemon-reload
systemctl restart docker
```
```
docker login nexus.jasper.cn.com:8082
docker push nexus.jasper.com.cn:8082/xieyuxuan/nginxvue:v1.0.0
```