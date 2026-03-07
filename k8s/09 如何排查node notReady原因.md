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
# 检查节点上的核心服务
sudo systemctl status kubelet
```
```
journal -u kubectl //查看问题
```