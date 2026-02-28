### 查看not ready节点
```
kubectl get nodes
```
### 查看node日志
```
kubectl describe node k8s-node1
```
### 到问题节点查看问题
```
systemctl status kubectl 
```
```
journal -u kubectl //查看问题
```