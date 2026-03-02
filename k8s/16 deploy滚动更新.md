## 触发滚动更新(无状态应用)
```
注意直接修改本地文件不会触发滚动更新
kubectl edit deploy deploy名字
```
```
查看滚动更新过程
kubectl rollout status deploy deploy名字
```
```
修改单个值
deploy删除后就没有了
kubectl set image 资源类型/资源名称 容器名=镜像:版本
kubectl set image deployment/名字 nginx=nginx:1.7.9 (--record 后面查history的时候会显现)
```
## 回滚历史版本
```
kubectl rollout history deploy/nginx-deploy
```
```
查看详细修改信息
kubectl rollout history deploy/nginx-deploy --revision=2//查看版本2详细修改信息
```
```
kubectl rollout undo deployment/nginx-deploy --to-version=2//回滚到版本2
```
```
查看状态
kubectl rollout status deployment/nginx-deploy 
```
## 注意必须要有revisionHistoryLimit不为零才能回退
```
revisionHistoryLimit默认为10
```