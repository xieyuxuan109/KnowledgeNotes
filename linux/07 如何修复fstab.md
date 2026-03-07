# fstab被覆盖失效
```
# 先确保根目录可写
sudo mount -o remount,rw /

# 查看所有分区
lsblk -f
```