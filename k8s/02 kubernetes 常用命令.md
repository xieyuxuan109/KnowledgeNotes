# 常用命令
### 集群信息
```
# kubectl cluster-info
Kubernetes control plane is running at https://lb.kubesphere.local:6443
CoreDNS is running at https://lb.kubesphere.local:6443/api/v1/namespaces/kube-system/services/coredns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```
### 获取节点基本信息
```
# kubectl get nodes
NAME     STATUS   ROLES           AGE     VERSION
master   Ready    control-plane   7d13h   v1.32.5
node1    Ready    worker          7d13h   v1.32.5
node2    Ready    worker          7d13h   v1.32.5
```
### 获取节点扩展信息
```
# kubectl get nodes -o wide
NAME     STATUS   ROLES           AGE     VERSION   INTERNAL-IP     EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION       CONTAINER-RUNTIME
master   Ready    control-plane   7d13h   v1.32.5   10.17.200.201   <none>        Ubuntu 22.04.5 LTS   5.15.0-163-generic   containerd://1.7.13
node1    Ready    worker          7d13h   v1.32.5   10.17.200.202   <none>        Ubuntu 22.04.5 LTS   5.15.0-163-generic   containerd://1.7.28
node2    Ready    worker          7d13h   v1.32.5   10.17.200.203   <none>        Ubuntu 22.04.5 LTS   5.15.0-163-generic   containerd://1.7.28
```
### 获取节点标签信息
```
# kubectl get nodes --show-labels
NAME     STATUS   ROLES           AGE     VERSION   LABELS
master   Ready    control-plane   7d13h   v1.32.5   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=master,kubernetes.io/os=linux,node-role.kubernetes.io/control-plane=,node.kubernetes.io/exclude-from-external-load-balancers=
node1    Ready    worker          7d13h   v1.32.5   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=node1,kubernetes.io/os=linux,node-role.kubernetes.io/worker=
node2    Ready    worker          7d13h   v1.32.5   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=node2,kubernetes.io/os=linux,node-role.kubernetes.io/worker=
```
### 获取节点详细信息
```
# kubectl describe node master
Name:               master
Roles:              control-plane
Labels:             beta.kubernetes.io/arch=amd64
                    beta.kubernetes.io/os=linux
                    kubernetes.io/arch=amd64
                    kubernetes.io/hostname=master
                    kubernetes.io/os=linux
                    node-role.kubernetes.io/control-plane=
                    node.kubernetes.io/exclude-from-external-load-balancers=
Annotations:        kubeadm.alpha.kubernetes.io/cri-socket: unix:///run/containerd/containerd.sock
                    node.alpha.kubernetes.io/ttl: 0
                    projectcalico.org/IPv4Address: 10.17.200.201/16
                    projectcalico.org/IPv4IPIPTunnelAddr: 11.233.70.0
                    volumes.kubernetes.io/controller-managed-attach-detach: true
CreationTimestamp:  Sat, 06 Dec 2025 13:02:33 +0000
Taints:             node-role.kubernetes.io/control-plane:NoSchedule
Unschedulable:      false
Lease:
  HolderIdentity:  master
  AcquireTime:     <unset>
  RenewTime:       Sun, 14 Dec 2025 02:49:00 +0000
Conditions:
  Type                 Status  LastHeartbeatTime                 LastTransitionTime                Reason                       Message
  ----                 ------  -----------------                 ------------------                ------                       -------
  NetworkUnavailable   False   Sun, 14 Dec 2025 01:08:47 +0000   Sun, 14 Dec 2025 01:08:47 +0000   CalicoIsUp                   Calico is running on this node
  MemoryPressure       False   Sun, 14 Dec 2025 02:45:29 +0000   Sat, 06 Dec 2025 13:02:26 +0000   KubeletHasSufficientMemory   kubelet has sufficient memory available
  DiskPressure         False   Sun, 14 Dec 2025 02:45:29 +0000   Sat, 06 Dec 2025 13:02:26 +0000   KubeletHasNoDiskPressure     kubelet has no disk pressure
  PIDPressure          False   Sun, 14 Dec 2025 02:45:29 +0000   Sat, 06 Dec 2025 13:02:26 +0000   KubeletHasSufficientPID      kubelet has sufficient PID available
  Ready                True    Sun, 14 Dec 2025 02:45:29 +0000   Sat, 06 Dec 2025 13:27:36 +0000   KubeletReady                 kubelet is posting ready status
Addresses:
  InternalIP:  10.17.200.201
  Hostname:    master
Capacity:
  cpu:                2
  ephemeral-storage:  24590672Ki
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  memory:             3968820Ki
  pods:               110
Allocatable:
  cpu:                1600m
  ephemeral-storage:  24590672Ki
  hugepages-1Gi:      0
  hugepages-2Mi:      0
  memory:             3336580093
  pods:               110
System Info:
  Machine ID:                 325923b3b7dc4970858ade937ce440a0
  System UUID:                79224d56-8e55-c7e6-da39-1cd8356504e8
  Boot ID:                    d275e938-b006-4e68-a816-69d44f4f7e95
  Kernel Version:             5.15.0-163-generic
  OS Image:                   Ubuntu 22.04.5 LTS
  Operating System:           linux
  Architecture:               amd64
  Container Runtime Version:  containerd://1.7.13
  Kubelet Version:            v1.32.5
  Kube-Proxy Version:         v1.32.5
PodCIDR:                      11.233.64.0/24
PodCIDRs:                     11.233.64.0/24
Non-terminated Pods:          (10 in total)
  Namespace                   Name                                        CPU Requests  CPU Limits  Memory Requests  Memory Limits  Age
  ---------                   ----                                        ------------  ----------  ---------------  -------------  ---
  default                     node-exporter-fz9jw                         100m (6%)     200m (12%)  200Mi (6%)       200Mi (6%)     13h
  kube-system                 calico-kube-controllers-678fc69664-l48p5    0 (0%)        0 (0%)      0 (0%)           0 (0%)         7d13h
  kube-system                 calico-node-zkqd4                           250m (15%)    0 (0%)      0 (0%)           0 (0%)         7d13h
  kube-system                 coredns-7ddb55578d-hc6h8                    100m (6%)     0 (0%)      70Mi (2%)        300Mi (9%)     7d13h
  kube-system                 coredns-7ddb55578d-hxfbk                    100m (6%)     0 (0%)      70Mi (2%)        300Mi (9%)     7d13h
  kube-system                 kube-apiserver-master                       250m (15%)    0 (0%)      0 (0%)           0 (0%)         7d13h
  kube-system                 kube-controller-manager-master              200m (12%)    0 (0%)      0 (0%)           0 (0%)         7d13h
  kube-system                 kube-proxy-qfshx                            0 (0%)        0 (0%)      0 (0%)           0 (0%)         7d13h
  kube-system                 kube-scheduler-master                       100m (6%)     0 (0%)      0 (0%)           0 (0%)         7d13h
  kube-system                 nodelocaldns-g2jgt                          100m (6%)     0 (0%)      70Mi (2%)        200Mi (6%)     7d13h
Allocated resources:
  (Total limits may be over 100 percent, i.e., overcommitted.)
  Resource           Requests     Limits
  --------           --------     ------
  cpu                1200m (75%)  200m (12%)
  memory             410Mi (12%)  1000Mi (31%)
  ephemeral-storage  0 (0%)       0 (0%)
  hugepages-1Gi      0 (0%)       0 (0%)
  hugepages-2Mi      0 (0%)       0 (0%)
Events:              <none>
```