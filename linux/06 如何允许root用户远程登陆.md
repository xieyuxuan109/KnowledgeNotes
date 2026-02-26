### 如何允许root用户远程登录
#### 修改root用户密码
```
sudo passwd root
```
#### 编辑文件
```
cd /etc/ssh
ls
vi sshd_config//找到PermitRootLogin 注释掉这一行
//添加PermitRootLogin yes 这一行
```
#### 重启生效
```
systemctl restart ssh
``` 
#### 查看
```
ps -au|grep sshd//有root用户说明设置成功
```