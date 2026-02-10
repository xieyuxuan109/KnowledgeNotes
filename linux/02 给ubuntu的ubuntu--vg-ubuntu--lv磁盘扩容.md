### 磁盘扩容
#### 查看空间
```
# df -lh
文件系统                           容量  已用  可用 已用% 挂载点
udev                               1.9G     0  1.9G    0% /dev
tmpfs                              389M  1.9M  387M    1% /run
/dev/mapper/ubuntu--vg-ubuntu--lv   24G   22G  914M   96% /
tmpfs                              1.9G   21M  1.9G    2% /dev/shm
tmpfs                              5.0M  4.0K  5.0M    1% /run/lock
tmpfs                              1.9G     0  1.9G    0% /sys/fs/cgroup
/dev/loop1                          92M   92M     0  100% /snap/lxd/36554
/dev/loop0                          64M   64M     0  100% /snap/core20/1828
/dev/loop2                          64M   64M     0  100% /snap/core20/2682
/dev/sda2                          2.0G  119M  1.7G    7% /boot
/dev/loop4                          51M   51M     0  100% /snap/snapd/25577
/dev/loop3                          50M   50M     0  100% /snap/snapd/18357
/dev/loop5                          92M   92M     0  100% /snap/lxd/24061
tmpfs                              389M   32K  389M    1% /run/user/1000

# fdisk -l
Disk /dev/loop0：63.29 MiB，66359296 字节，129608 个扇区
单元：扇区 / 1 * 512 = 512 字节
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节


Disk /dev/loop1：91.92 MiB，96358400 字节，188200 个扇区
单元：扇区 / 1 * 512 = 512 字节
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节


Disk /dev/loop2：63.79 MiB，66871296 字节，130608 个扇区
单元：扇区 / 1 * 512 = 512 字节
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节


Disk /dev/loop3：49.86 MiB，52260864 字节，102072 个扇区
单元：扇区 / 1 * 512 = 512 字节
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节


Disk /dev/loop4：50.95 MiB，53399552 字节，104296 个扇区
单元：扇区 / 1 * 512 = 512 字节
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节


Disk /dev/loop5：91.85 MiB，96292864 字节，188072 个扇区
单元：扇区 / 1 * 512 = 512 字节
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节




Disk /dev/sda：50 GiB，53687091200 字节，104857600 个扇区
Disk model: VMware Virtual S
单元：扇区 / 1 * 512 = 512 字节
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节
磁盘标签类型：gpt
磁盘标识符：C9A53B1D-7F40-40CC-9BAD-5D35D127F4BC

设备          起点      末尾      扇区 大小 类型
/dev/sda1     2048      4095      2048   1M BIOS 启动
/dev/sda2     4096   4198399   4194304   2G Linux 文件系统
/dev/sda3  4198400 104855551 100657152  48G Linux 文件系统


Disk /dev/mapper/ubuntu--vg-ubuntu--lv：23.102 GiB，25765609472 字节，50323456 个扇区
单元：扇区 / 1 * 512 = 512 字节
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节
```
#### 命令vgdisplay查看lvm卷组的信息，如果看到 Free PE / Size > 0，表示还有扩容空间。
```
# vgdisplay
  --- Volume group ---
  VG Name               ubuntu-vg
  System ID             
  Format                lvm2
  Metadata Areas        1
  Metadata Sequence No  2
  VG Access             read/write
  VG Status             resizable
  MAX LV                0
  Cur LV                1
  Open LV               1
  Max PV                0
  Cur PV                1
  Act PV                1
  VG Size               <48.00 GiB
  PE Size               4.00 MiB
  Total PE              12287
  Alloc PE / Size       6143 / <24.00 GiB
  Free  PE / Size       6144 / 24.00 GiB
  VG UUID               63amyH-FEmO-bAst-Yhck-o1Fo-qBkv-QgyxGb
```
#### 使用lvresize扩容剩下所有空间
```
#lvextend -L 10G /dev/mapper/ubuntu--vg-ubuntu--lv      //增大或减小至19G
#lvextend -L +10G /dev/mapper/ubuntu--vg-ubuntu--lv     //增加10G
#lvreduce -L -10G /dev/mapper/ubuntu--vg-ubuntu--lv     //减小10G
#lvresize -l  +100%FREE /dev/mapper/ubuntu--vg-ubuntu--lv   //按百分比扩
```
#### 扩容
```
# lvresize -l +100%FREE /dev/mapper/ubuntu--vg-ubuntu--lv
  Size of logical volume ubuntu-vg/ubuntu-lv changed from <24.00 GiB (6143 extents) to <48.00 GiB (12287 extents).
  Logical volume ubuntu-vg/ubuntu-lv successfully resized.
```
#### 使扩容生效
```
# resize2fs /dev/mapper/ubuntu--vg-ubuntu--lv
resize2fs 1.45.5 (07-Jan-2020)
/dev/mapper/ubuntu--vg-ubuntu--lv 上的文件系统已被挂载于 /；需要进行在线调整大小
old_desc_blocks = 3, new_desc_blocks = 6
/dev/mapper/ubuntu--vg-ubuntu--lv 上的文件系统大小已经调整为 12581888 个块（每块 4k）。
```
#### 再次查看
```
# vgdisplay
--- Volume group --- VG Name ubuntu-vg System ID
Format lvm2 Metadata Areas 1 Metadata Sequence No 3 VG Access read/write VG Status resizable MAX LV 0 Cur LV 1 Open LV 1 Max PV 0 Cur PV 1 Act PV 1 VG Size <48.00 GiB PE Size 4.00 MiB Total PE 12287 Alloc PE / Size 12287 / <48.00 GiB Free PE / Size 0 / 0
VG UUID 63amyH-FEmO-bAst-Yhck-o1Fo-qBkv-QgyxGb
```