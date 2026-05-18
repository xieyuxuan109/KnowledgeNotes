```
配置
vim /etc/hosts
10.17.210.200 nexus.dennis.com.cn
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
{
  "insecure-registries": ["nexus.dennis.com.cn:8082","nexus.dennis.com.cn:8083"],
  "registry-mirrors": [
    "http://nexus.dennis.com.cn:8083"
  ]
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