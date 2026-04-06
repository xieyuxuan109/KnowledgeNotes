### 1. 立即清理 Docker（最快释放空间）
```
# 清理所有未使用的 Docker 资源（包括停止的容器、未使用的镜像、网络、构建缓存）
docker system prune -a -f --volumes
# 查看释放了多少空间
docker system df
```
### 3. 清理系统日志
```
# 清理 systemd 日志（限制为 200MB）
sudo journalctl --vacuum-size=200M
# 清理旧的日志文件
sudo find /var/log -type f -name "*.log" -mtime +30 -delete
sudo find /var/log -type f -name "*.gz" -delete
```
### 4. 清理 apt 缓存（Ubuntu/Debian）
```
# 清理软件包缓存
sudo apt-get clean
sudo apt-get autoremove -y
```
### 5. 检查并清理大文件
```
# 查看 /tmp 目录
sudo du -sh /tmp/* 2>/dev/null | sort -rh | head -5
# 清理临时文件
sudo rm -rf /tmp/*
```
### 清理后检查空间
```
df -h
```