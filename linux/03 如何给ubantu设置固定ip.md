### 如何给ubantu设置固定ip
#### 查看网关地址和需要设置的网卡
```
root@jasper:/home/jasper# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.1.1     0.0.0.0         UG    0      0        0 ens33
172.17.0.0      0.0.0.0         255.255.0.0     U     0      0        0 docker0
172.23.0.0      0.0.0.0         255.255.0.0     U     0      0        0 br-3b986a381cfa
192.168.1.0     0.0.0.0         255.255.255.0   U     0      0        0 ens33

```
#### 编辑文件
```
sudo vim /etc/netplan/50-cloud-init.yaml
```
```
root@jasper:/home/jasper# cat /etc/netplan/50-cloud-init.yaml
# This is the network config written by 'subiquity'
network:
  ethernets:
    ens33: 
      addresses: [10.17.220.221/16]
      routes:
        - to: default
          via: 10.17.0.1
      dhcp4: false
      nameservers:
        addresses: [114.114.114.114,8.8.8.8,10.17.0.1]
  version: 2
```

#### 输入命令使改动生效
```
netplan apply
```