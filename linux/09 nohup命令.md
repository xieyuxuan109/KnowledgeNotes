## 后台执行脚本
```
nohup /root/runoob.sh &
```
## 查找运行的脚本PID
```
ps -aux | grep "runoob.sh" 
```
## 相关参数说明
```
参数说明：

a : 显示所有程序
u : 以用户为主的格式来显示
x : 显示所有程序，不区分终端机
```
## 删除进程
```
kill -9  进程号PID
```
## 错误输出
```
nohup /root/runoob.sh > runoob.log 2>&1 & //定向输入到文件
```
```
2>&1 解释：

将标准错误 2 重定向到标准输出 &1 ，标准输出 &1 再被重定向输入到 runoob.log 文件中。

0 – stdin (standard input，标准输入)
1 – stdout (standard output，标准输出)
2 – stderr (standard error，标准错误输出)
```
## 第二个安装screen命令
```
sudo apt install screen
```
## 基本使用
```
启动新会话：输入 screen 或 screen -S 会话名 创建新会话。
查看会话列表：使用 screen -ls 查看当前运行的所有会话。
恢复会话：通过 screen -r 会话名 恢复已断开的会话。
退出会话：在会话中输入 exit 关闭当前会话。
```
## 常用快捷键
```
Ctrl + A + D：将当前会话挂起并返回主终端。
Ctrl + A + C：创建新窗口。
Ctrl + A + N：切换到下一个窗口。
Ctrl + A + K：关闭当前窗口。
```