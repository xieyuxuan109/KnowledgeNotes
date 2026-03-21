## 安装helm
```
root@k8s-master:/home/jasper/helm# curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 11929  100 11929    0     0  10080      0  0:00:01  0:00:01 --:--:-- 10083
Downloading https://get.helm.sh/helm-v3.20.1-linux-amd64.tar.gz
Verifying checksum... Done.
Preparing to install helm into /usr/local/bin
helm installed into /usr/local/bin/helm

root@k8s-master:/home/jasper/helm# helm version
version.BuildInfo{Version:"v3.20.1", GitCommit:"a2369ca71c0ef633bf6e4fccd66d634eb379b371", GitTreeState:"clean", GoVersion:"go1.25.8"}
```
## 添加helm仓库
```
# 添加仓库
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

# 查看仓库列表
helm repo list

# 搜索 ingress-nginx
helm search repo ingress-nginx
# 下载安装包
helm pull ingress-nginx/ingress-nginx
# 解压
tar -xf ingress-nginx-4.15.1.tgz
```