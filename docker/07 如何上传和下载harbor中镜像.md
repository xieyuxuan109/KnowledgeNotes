### 上传镜像
#### 修改配置
```
root@jasper:/home/jasper/harbor# cat /etc/docker/daemon.json 
{
    "proxies":{
        "http-proxy":"http://192.168.1.102:8118",
        "https-proxy":"http://192.168.1.102:8118"
    },
  "insecure-registries" : ["192.168.1.221:81"]
}
root@jasper:/home/jasper/harbor# systemctl daemon-reload
root@jasper:/home/jasper/harbor# systemctl docker restart
```
#### 登录
```
# docker login ip:prot -u xxx -p xxx

docker login 192.168.182.110:80 -u admin
Harbor12345
```
#### 进行tag
```
# docker tag image_id(本地需要push的镜像) ip:port/项目名/保存的镜像名(例如：xxx:version1)
docker tag goharbor/harbor-exporter:v2.5.6 local-168-182-110:80/library/goharbor/harbor-exporter:v2.5.6
```
#### 推送
```
# docker push ip:port/项目名/保存的镜像名(例如：xxx:version1)
docker push local-168-182-110:80/library/goharbor/harbor-exporter:v2.5.6
```
### 拉取镜像
```
# docker pull ip:port/项目名/保存的镜像名(例如：xxx:version1)
docker pull local-168-182-110:80/library/goharbor/harbor-exporter:v2.5.6
```