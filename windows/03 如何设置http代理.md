### privoxy进行http代理
#### windows不区分大小写上下两组挑一组即可
#### linux最好全部都写
```
export HTTP_PROXY=http://10.17.200.100:8118
export HTTPS_PROXY=http://10.17.200.100:8118
export http_proxy=http://10.17.200.100:8118
export https_proxy=http://10.17.200.100:8118
curl www.google.com
```
```
unset HTTP_PROXY
unset HTTPS_PROXY
unset http_proxy
unset https_proxy
```
#### docker需要另外进行代理配置
```
sudo vim /etc/docker/daemon.json
```
```
{
    "proxies":{
        "http-proxy":"http://10.17.200.100:8118",
        "https-proxy":"http://10.17.200.100:8118"
    }
}
```
#### 重启生效 
```
root@jasper:/home/jasper# systemctl daemon-reload
root@jasper:/home/jasper# systemctl restart docker
```