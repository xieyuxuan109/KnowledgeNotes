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
## nexus自动登录配置文件
```
//~当前用户的根目录下
root@dennis:/home/dennis/Blogx# ls ~/ -a
.                        .cache              jenkins-5.9.12.tgz                          .python_history
..                       cert-manager.yaml   .kube                                       review.py
base-1.29.1.tgz          cilium-1.19.2.tgz   .lesshst                                    snap
.bash_history            .config             .local                                      .ssh
.bash_history-16148.tmp  coredns-1.45.2.tgz  nfs-subdir-external-provisioner-4.0.18.tgz  test.sh
.bash_history-31884.tmp  .docker             opentelemetry-collector-0.147.2.tgz         .vim
.bash_history-32964.tmp  elk                 opentelemetry-operator-0.109.0.tgz          .viminfo
.bash_history-44393.tmp  istiod-1.29.1.tgz   proc                                        .Xauthority
.bashrc                  jenkins             .profile
root@dennis:/home/dennis/Blogx# ls ~/.docker/
buildx  config.json
root@dennis:/home/dennis/Blogx# cat ~/.docker/config.json 
{
	"auths": {
		"192.168.1.200:8082": {
			"auth": "ZGVubmlzOnhpZWNodWFuMTEwJg=="
		},
		"nexus.dennis.com.cn:8082": {
			"auth": "YWRtaW46eGllY2h1YW4xMTAm"
		},
		"nexus.dennis.com.cn:8083": {
			"auth": "YWRtaW46eGllY2h1YW4xMTAm"
		}
	}
}
```