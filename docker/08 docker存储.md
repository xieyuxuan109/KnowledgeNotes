### docker存储区分
### 目录挂载：-v /app/nghtml:/usr/share/nginx/html 以主机为准，如果主机中目录为空，挂载后docker容器中也为空
### 卷映射：-v ngconf:/etc/nginx 以容器为准，容器中目录里面是什么外面卷映射是什么
#### 卷映射默认存储路径 /var/lib/docker/volumes/<volume-name>
#### 删除容器时候，挂载的目录和卷映射都不会被删除
### 卷映射常用命令 
```
root@jasper:/home/jasper# docker volume ls
DRIVER    VOLUME NAME
root@jasper:/home/jasper# docker volume create haha
haha
root@jasper:/home/jasper# docker volume ls
DRIVER    VOLUME NAME
local     haha
root@jasper:/home/jasper# docker volume inspect haha
[
    {
        "CreatedAt": "2026-02-13T11:40:34Z",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/lib/docker/volumes/haha/_data",
        "Name": "haha",
        "Options": null,
        "Scope": "local"
    }
]
root@jasper:/home/jasper# docker volume rm haha
haha
root@jasper:/home/jasper# docker volume  ls
DRIVER    VOLUME NAME

```