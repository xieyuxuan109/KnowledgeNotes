# Kubernetes Pod 通用模板
## 简单创建pod
```
kubectl create pod 名字 --image=镜像名称
pod名称：Deployment名 - RS哈希 - Pod随机串
```
## 获取基本配置文件
```
kubectl get pod 名字 -o yaml
```
```
在 Kubernetes 资源清单里，apiVersion 表示 这个资源对象属于哪个 API 组和版本。
apiVersion: v1
apiVersion: apps/v1
它们的区别主要在于：是否属于某个 API 组（group）。
一、apiVersion: v1
v1 表示它属于 核心 API 组（core group）。
核心组是 Kubernetes 最基础的一组资源，没有 group 名字，所以直接写 v1。
常见属于 v1 的资源：
Pod
Service
ConfigMap
Secret
Namespace
Node
PersistentVolume
PersistentVolumeClaim
例如：
apiVersion: v1
kind: Pod
二、apiVersion: apps/v1
apps/v1 表示：
API 组：apps
版本：v1
这是扩展 API 组的一部分，用来管理控制器类型资源。
属于 apps/v1 的资源：
Deployment
ReplicaSet
StatefulSet
DaemonSet
例如：
apiVersion: apps/v1
kind: Deployment
```
```yaml
#示例
apiVersion: v1 #api文档版本
kind: Pod #资源对象类型，也可以配置为Deployment StatefulSet 等对象
metadata: #Pod相关元数据，用于描述Pod的数据
  name: nginx-demo #Pod名称
  labels: #定义Pod标签
    type: app #label自己随便写
    version: 1.0.0 #自定义label标签 名字为type 值为app
  namespace: 'default' #默认default
spec: #期望Pod按照这里面的描述进行创建
  containers:
  - name: nginx # 容器的名称
    image: nginx # 指定容器的镜像
    imagePullPolicy: IfNotPresent # Always IfNotPresent Never三种策略 如果镜像标签是 latest：默认值是 Always，表示每次启动容器时都会拉取镜像。如果镜像标签不是 latest：默认值是 IfNotPresent，表示只有本地没有该镜像时才会拉取。
    startupProbe: #应用启动探针配置
      #httpGet: #探测方式
      #path: /index.html #http探测路径
      #tcpSocket:
      exec:
        command:
        - sh
        - -c
        - "sleep 1;echo 'success'>/inited;"
          #port: 80  #请求端口
      timeoutSeconds: 10 #超时时间
      successThreshold: 1 #检测1次成功就表示成功
      failureThreshold: 2 # 检测失败两次就表示失败
      periodSeconds: 10 #间隔时间
    command: # 指定容器启动时执行的命令
      - nginx
      - -g
      - 'daemon off;' # nginx -g 'daemon off;'
    workingDir: /usr/share/nginx/html # 定义容器启动后的目录
    ports:
    - name: http
      containerPort: 80 # 容器内暴露什么端口
      protocol: TCP #描述端口是基于哪种协议通信的
    env: #环境变量
    - name: JVM_OPTS # 环境变量名称
      value: '-Xms128m -Xmx128m' # 环境变量的值，只是简单的描述
    resources:
      requests: # 最少需要多少资源
        cpu: 100m # 限制cpu最少使用多少个 1000m等于一个核心 100m等于0.1个核心
        memory: 128Mi # 限制内存最少使用128M
      limits: # 最多可以用多少资源
        cpu: 200m # 限制cpu最少使用0.2个核心
        memory: 256Mi # 限制最多使用256M
  restartPolicy: OnFailure # 重启策略，OnFailure 只有失败才会重启 Always pod一旦终止运行，则无论容器如何终止，kubelet服务都将重启它
  # imagePullSecret 私有仓库需要登陆拉取时候配置用户名密码
```
---

# 常见字段解释

## 一、基础字段

| 字段 | 说明 |
|------|------|
| apiVersion | API 版本，Pod 固定为 v1 |
| kind | 资源类型 |
| metadata | 元数据 |
| spec | 具体规格定义 |

---

## 二、metadata 常见字段

| 字段 | 作用 |
|------|------|
| name | Pod 名称 |
| namespace | 所属命名空间 |
| labels | 标签（用于选择器匹配） |
| annotations | 注解（存储说明信息，不参与选择） |

---

## 三、spec 常见字段

### restartPolicy

| 值 | 说明 |
|----|------|
| Always | 默认，失败自动重启 |
| OnFailure | 失败才重启 |
| Never | 不重启 |

---

### containers（核心字段）

一个 Pod 至少包含一个 container。

| 字段 | 作用 |
|------|------|
| name | 容器名称 |
| image | 镜像 |
| imagePullPolicy | 拉取策略 |
| ports | 容器端口 |
| env | 环境变量 |
| resources | 资源请求与限制 |
| volumeMounts | 挂载卷 |
| securityContext | 安全设置 |

---

### resources

- requests：调度时保证的资源
- limits：最大可用资源

示例：
```yaml
resources:
  requests:
    cpu: 100m
    memory: 128Mi
  limits:
    cpu: 500m
    memory: 256Mi
```

---

### probes（健康检查）

- livenessProbe：存活探针（失败会重启）
- readinessProbe：就绪探针（失败会从 Service 摘除）

支持：
- httpGet
- tcpSocket
- exec

---

### volumes

常见类型：

| 类型 | 说明 |
|------|------|
| emptyDir | 临时目录 |
| hostPath | 宿主机目录 |
| configMap | 配置 |
| secret | 密钥 |
| persistentVolumeClaim | 持久卷 |

---

### 调度相关

| 字段 | 作用 |
|------|------|
| nodeSelector | 指定节点标签 |
| tolerations | 容忍污点 |
| affinity | 亲和性调度 |

---

# 面试高频考点总结

1. requests 决定调度，limits 决定最大使用
2. liveness 失败会重启容器
3. readiness 失败不会重启，只是摘除流量
4. Pod 本身不做负载均衡，需要 Service
5. Pod 不建议直接创建生产使用，通常由 Deployment 管理