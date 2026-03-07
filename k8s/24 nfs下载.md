## nfs下载
```
sudo apt install nfs-kernel-server -y
sudo systemctl status nfs-kernel-server
sudo mkdir -p /mnt/nfs_share
sudo chmod 777 /mnt/nfs_share
sudo nano /etc/exports
```
```
# 格式: 共享目录 允许访问的IP(权限选项)

# 示例1: 共享给单个客户端 (IP为 192.168.1.100)
/mnt/nfs_share 192.168.1.100(rw,sync,no_subtree_check)

# 示例2: 共享给整个子网 (IP段 192.168.1.0/24)
/mnt/nfs_share 192.168.1.0/24(rw,sync,no_subtree_check)

# 示例3: 共享给所有客户端 (不推荐用于生产环境)
/mnt/nfs_share *(rw,sync,no_subtree_check)
```
```
# 导出所有共享目录
sudo exportfs -a

# 重启NFS服务使配置生效
sudo systemctl restart nfs-kernel-server
```
## 配置NFS客户端
```
sudo apt install nfs-common -y
sudo mkdir -p /mnt/nfs_client_share
# 格式: sudo mount 服务器IP:共享目录 本地挂载点
sudo mount 192.168.1.100:/mnt/nfs_share /mnt/nfs_client_share
```
## 测试
```
在服务器的共享目录中创建一个测试文件：

bash
sudo touch /mnt/nfs_share/testfile.txt
```
```
在客户端的挂载点目录中查看：

bash
ls /mnt/nfs_client_share
```