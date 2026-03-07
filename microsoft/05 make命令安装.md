# windows make命令安装 使用powershell
## 安装 Scoop（只需要一次）
```
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
irm get.scoop.sh | iex
```
## 安装 make
```
scoop install make
```
## 测试
```
make -v
```
```
安装完成后：
PowerShell 可以用
CMD 也可以用
Git Bash 也可以用
因为 Scoop 会自动把路径加入 Windows 的 PATH 环境变量。
```