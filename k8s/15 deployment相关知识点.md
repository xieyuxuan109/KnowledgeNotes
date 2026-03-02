## 简单创建deploy 
```
kubectl create deploy 名字 --image=镜像名称 --replicas=3
```
## 获取基本配置文件
```
kubectl get deploy 名字 -o yaml
```
## 查看
```
kubectl get deploy(/deployments) (-o wide详细信息)
deploy会帮我们创建对应的replicaset和pod
kubectl get rs(/replicaset) (-o wide详细信息)
kubectl get po(/pod) (-o wide详细信息)
```
## 模板
```yaml
apiVersion: apps/v1        # deployment api 版本
kind: Deployment           # 资源类型为 deployment

metadata:                  # 元信息
  labels:                  # 标签
    app: nginx-deploy      # 具体的 key: value 配置形式
  name: nginx-deploy       # deployment 的名字
  namespace: default       # 所在的命名空间

spec:
  replicas: 1              # 期望副本数
  revisionHistoryLimit: 10 # 进行滚动更新后，保留的历史版本数

  selector:                # 选择器，用于找到匹配的 RS
    matchLabels:           # 按照标签匹配
      app: nginx-deploy    # 匹配的标签 key/value

  strategy:                # 更新策略
    rollingUpdate:         # 滚动更新配置
      maxSurge: 25%        # 更新过程中 最多可以比期望副本数多创建多少新 Pod
      maxUnavailable: 25%  # 更新过程中，最多允许多少个 Pod 不可用
    type: RollingUpdate    # 更新类型，采用滚动更新

  template:                # pod 模板
    metadata:              # pod 的元信息
      labels:              # pod 的标签
        app: nginx-deploy

    spec:                  # pod 规格信息
      containers:          # pod 的容器
        - image: nginx:1.7.9          # 镜像
          imagePullPolicy: IfNotPresent # 拉取策略
          name: nginx                 # 容器名称

      restartPolicy: Always           # 重启策略
      terminationGracePeriodSeconds: 30 # 删除操作最大宽限时间
```
