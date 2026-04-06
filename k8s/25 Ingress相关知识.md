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
```
root@k8s-master:/home/jasper/helm# ls
ingress-nginx  ingress-nginx-4.15.1.tgz
root@k8s-master:/home/jasper/helm# cd ingress-nginx
root@k8s-master:/home/jasper/helm/ingress-nginx# ls
changelog  Chart.yaml  ci  cloudbuild.yaml  OWNERS  README.md  README.md.gotmpl  templates  tests  values.yaml
root@k8s-master:/home/jasper/helm/ingress-nginx# vim values.yaml 
//注释#digest: sha256:01038e7de14b78d702d2849c3aad72fd25903c4765af63cf16aa3398f5d5f2dd
dnsPolicy: ClusterFirstWithHostNet
hostNetWork: true
kind:deployment->Deamonset
node-selector->ingress: "true"
LoadBalancer->type: ClusterIP
enabled: false
```
```
安装
root@k8s-master:/home/jasper/helm/ingress-nginx# kubectl create ns ingress-nginx
namespace/ingress-nginx created
root@k8s-master:/home/jasper/helm/ingress-nginx# kubectl label node k8s-master ingress=true 
node/k8s-master labeled
root@k8s-master:/home/jasper/helm/ingress-nginx# helm install ingress-nginx -n ingress-nginx .
```
```
NAME: ingress-nginx
LAST DEPLOYED: Sun Mar 29 09:03:24 2026
NAMESPACE: ingress-nginx
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
The ingress-nginx controller has been installed.
Get the application URL by running these commands:
  export POD_NAME="$(kubectl get pods --namespace ingress-nginx --selector app.kubernetes.io/name=ingress-nginx,app.kubernetes.io/instance=ingress-nginx,app.kubernetes.io/component=controller --output jsonpath="{.items[0].metadata.name}")"
  kubectl port-forward --namespace ingress-nginx "${POD_NAME}" 8080:80
  echo "Visit http://127.0.0.1:8080 to access your application."

An example Ingress that makes use of the controller:
  apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    name: example
    namespace: foo
  spec:
    ingressClassName: nginx
    rules:
      - host: www.example.com
        http:
          paths:
            - pathType: Prefix
              backend:
                service:
                  name: exampleService
                  port:
                    number: 80
              path: /
    # This section is only required if TLS is to be enabled for the Ingress
    tls:
      - hosts:
        - www.example.com
        secretName: example-tls

If TLS is enabled for the Ingress, a Secret containing the certificate and key must also be provided:

  apiVersion: v1
  kind: Secret
  metadata:
    name: example-tls
    namespace: foo
  data:
    tls.crt: <base64 encoded cert>
    tls.key: <base64 encoded key>
  type: kubernetes.io/tls
```
```
411  kubectl create ns ingress-nginx
  412  kubectl label node k8s-master ingress=true 
  413  helm install ingress-nginx -n ingress-nginx .
  414  kubectl get po -n ingress-nginx
  415  kubectl get nodes --show-labels
  416  kubectl label no k8s-node1 ingress=true
  417  kubectl get po -n ingress-nginx
```
```
apiVersion: networking.k8s.io/v1  # 使用 Kubernetes 网络 API 的稳定版本
kind: Ingress                      # 资源类型为 Ingress，用于定义外部访问集群服务的规则

metadata:
  name: jasper-ingress             # Ingress 资源名称，集群内唯一
  # 注意：不再使用 kubernetes.io/ingress.class 注解，改用 spec.ingressClassName

spec:
  ingressClassName: nginx          # 指定使用哪个 Ingress Controller（这里为 nginx）
                                   # 替代了旧版注解 kubernetes.io/ingress.class

  rules:                           # 路由规则列表，可以配置多个 host
  - host: k8s.wolfcode.cn          # 外部访问的域名（支持通配符如 *.wolfcode.cn）
    http:                          # HTTP 协议规则
      paths:                       # 路径匹配规则列表，相当于 nginx 的 location 配置
      - pathType: Prefix           # 路径匹配类型：
                                   #   - Prefix：前缀匹配，以 / 为分隔符
                                   #   - Exact：精确匹配，需完全一致且区分大小写
                                   #   - ImplementationSpecific：由 IngressClass 决定
        path: /api                 # 匹配的 URL 路径，这里表示所有以 /api 开头的请求
        backend:                   # 后端服务配置
          service:
            name: nginx-svc        # 要代理的目标 Service 名称
            port:
              number: 80           # Service 的端口号（也可简写为 port: 80）
```