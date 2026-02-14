### 如何下载haobor
#### 下载地址 优先推荐离线下载包
```
https://github.com/goharbor/harbor/releases/
```
#### 解压
```
tar xvf harbor-offline-installer-v....tgz
cd harbor
```
#### 修改配置
```
1、hostname改为本机hostname或者本机IP 2、把http的端口改为自己想要映射的端口， 3、去掉https 4、修改密码
cp harbor.yml.tmpl harbor.yml

vim harbor.yml

hostname: {自己服务器的ip 内网外网都可以}

# htp related config
http:
# port for htp, default is 80. If htps enabled, this port will redirect to htps port
port: {自定义端口}

# https related config
#https:
  # https port for harbor, default is 443
#  port: 443
  # The path of cert and key files for nginx
#  certificate: /your/certificate/path
#  private_key: /your/private/key/path
```
#### 安装harbor
```
# 查看
docker-compose ps

# 再次安装，就可以执行以下命令
# docker-compose up -d
# 或者执行下面这句
# docker-compose up -f docker-compose.yml -d

# 停止
# docker-compose down
```
#### 访问
```
访问：http://ip:port 账号/密码：admin/Harbor12345（默认）
```