## 扩缩容方法
```
kubectl edit 来修改
```
```
kubectl scale --replicas=3 资源 资源名字
kubectl scale --replicas=3 deploy nginx-deploy
```
## 暂停与恢复
```
暂停
kubectl rollout pause deploy nginx-deploy
恢复
kubectl rollout pause resume nginx-deploy
```