### 如何在每个节点上面使用kubectl 
#### 先要修改节点配置，允许以root用户远程登陆
#### 在master节点中将配置发送给目标节点
```
scp /etc/kubernetes/admin.conf root@目标ip:~/.bash_profile
```
#### 添加环境变量
```
echo "export KUBECONFIG=/etc/kubernetes/admin.conf">>~/.bash_profile
```
#### 生效
```
source ~/.bash_profile
```
#### 验证kubectl 
```
kubectl get nodes
```