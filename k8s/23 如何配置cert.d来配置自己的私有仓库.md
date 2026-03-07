```
vim /etc/hosts
10.17.210.200 nexus.jasper.cn.com:8082
先配置域名解析
```
```
mkdir -p /etc/containerd/certs.d/nexus.jasper.cn.com:8082
vim /etc/containerd/certs.d/nexus.jasper.cn.com:8082/hosts.toml
```
```
server = "http://nexus.jasper.com.cn:8082"

[host."http://nexus.jasper.com.cn:8082"]
  capabilities = ["pull", "resolve", "push"]
  skip_verify = true
```
```
检查
cat /etc/containerd/config.toml | grep config_path
config_path = "/etc/containerd/certs.d"//config_path必须是这样
```
```
重启
systemctl daemon-reload
systemctl restart containerd
```
```
crictl pull \
--creds admin:密码 \
nexus.jasper.com.cn:8082/xieyuxuan/nginxvue:v1.0.0
```