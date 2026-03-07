# 一、创建 StatefulSet

## 创建 Headless Service（必须）

```bash
kubectl apply -f service.yaml
```

示例：

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: nginx
spec:
  clusterIP: None
  selector:
    app: nginx
  ports:
  - port: 80
---
apiVersion: apps/v1           # API版本，StatefulSet必须是 apps/v1
kind: StatefulSet             # 资源类型
metadata:
  name: mysql                   # StatefulSet 名称
  namespace: default          # 命名空间

spec:
  serviceName: mysql        # 必须关联一个 Headless Service
  replicas: 3                 # Pod 副本数量
  selector:                   # 标签选择器（必须和pod标签一致）
    matchLabels:
      app: mysql
  template:                   # Pod模板
    metadata:
      labels:
        app: mysql            # Pod标签
    spec:
      containers:
      - name: mysql           # 容器名称
        image: mysql:8.0     # 容器镜像
        ports:
        - containerPort: 3306   # 容器端口
          name: http
        volumeMounts:         # 挂载存储
        - name: data
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:       # PVC模板（每个pod一个独立存储）
  - metadata:
      name: data              # 必须和 volumeMounts name一致
    spec:
      accessModes:
      - ReadWriteOnce         # 访问模式
      resources:
        requests:
          storage: 1Gi        # 存储大小

```

查看 Service：

```bash
kubectl get svc
kubectl describe svc nginx
```

---

## 创建 StatefulSet

```bash
kubectl apply -f statefulset.yaml
```

查看：

```bash
kubectl get statefulsets
kubectl get sts
```

查看详情：

```bash
kubectl describe statefulset web
```

---

# 二、查看运行状态

## 查看 Pod

```bash
kubectl get pods
kubectl get pods -o wide
```

有序名称示例：

```
web-0
web-1
web-2
```

---

## 查看 PVC（重点）

```bash
kubectl get pvc
```

示例：

```
www-web-0
www-web-1
www-web-2
```

---

## 查看 PV

```bash
kubectl get pv
```

---

## 查看 DNS 是否正常

进入 Pod：

```bash
kubectl exec -it web-0 -- /bin/sh
```

测试 DNS：

```bash
ping web-1.nginx.default.svc.cluster.local
```

---

# 三、扩容与缩容（最常用）

## 扩容

```bash
kubectl scale statefulset web --replicas=5
```

查看变化：

```bash
kubectl get pods
```

新增：

```
web-3
web-4
```

---

## 缩容

```bash
kubectl scale statefulset web --replicas=2
```

删除顺序：

```
web-4
web-3
web-2
```

⚠ 注意：

```bash
kubectl get pvc
```

PVC 不会删除！

---

# 四、更新镜像（滚动更新）

##  修改镜像

```bash
kubectl set image statefulset/web nginx=nginx:1.26
```

查看更新状态：

```bash
kubectl rollout status statefulset/web
```

查看历史版本：

```bash
kubectl rollout history statefulset/web
```

---

##  回滚

```bash
kubectl rollout undo statefulset/web
```

指定版本回滚：

```bash
kubectl rollout undo statefulset/web --to-revision=1
```

---

# 五、删除操作（高频）

## 删除 Pod

```bash
kubectl delete pod web-1
```

特点：

- 会自动重建
- 名字不变
- PVC 不变

---

## 删除 StatefulSet

```bash
kubectl delete statefulset web
```

默认行为：

- Pod 删除
- PVC 不删除

---

##  删除 PVC（手动清理）

```bash
kubectl delete pvc www-web-0
```

全部删除：

```bash
kubectl delete pvc -l app=nginx
```

---

# 六、强制删除场景

## 强制删除卡住的 Pod

```bash
kubectl delete pod web-1 --force --grace-period=0
```

---

##  保留 Pod 删除 StatefulSet

```bash
kubectl delete statefulset web --cascade=orphan
```

作用：

- 只删除 StatefulSet
- 保留 Pod

---

# 七、查看事件排错

```bash
kubectl describe pod web-0
```

查看事件：

```bash
kubectl get events
```

按时间排序：

```bash
kubectl get events --sort-by=.metadata.creationTimestamp
```

---

# 八、修改配置（编辑方式）

在线编辑：

```bash
kubectl edit statefulset web
```

修改后自动滚动更新。

---

# 九、常见问题排查命令

##  PVC Pending

```bash
kubectl describe pvc www-web-0
```

查看 StorageClass：

```bash
kubectl get sc
```

---

## Pod 一直 Pending

```bash
kubectl describe pod web-0
```

---

##  镜像拉取失败

```bash
kubectl describe pod web-0
```

---

## 查看日志

```bash
kubectl logs web-0
```

指定容器：

```bash
kubectl logs web-0 -c nginx
```

---

# 十、更新策略操作

查看更新策略：

```bash
kubectl get sts web -o yaml | grep updateStrategy -A 3
```

修改为 OnDelete：

```yaml
updateStrategy:
  type: OnDelete
```

效果：

- 不自动更新
- 必须手动删除 Pod 才会更新

---

# 十一、命令速查表（面试必背）

## 查看

```
kubectl get sts
kubectl get pods
kubectl get pvc
kubectl describe sts web
```

## 扩缩容

```
kubectl scale sts web --replicas=5
```

## 更新

```
kubectl set image sts/web nginx=nginx:1.26
kubectl rollout status sts/web
kubectl rollout undo sts/web
```

## 删除

```
kubectl delete pod web-1
kubectl delete sts web
kubectl delete pvc xxx
```

---

