## 拉取镜像
```
docker pull gitlab/gitlab-ce:19.1.0-ce.0
docker pull gitlab/gitlab-runner:v19.1.0
```
## 一些准备
```
sudo mkdir -p /srv/gitlab
export GITLAB_HOME=/srv/gitlab
docker volume create gitlab-runner-config
```
## docker compose
```
version: "3.9"

services:

  # GitLab 服务
  gitlab:
    # GitLab CE 社区版镜像
    image: nexus.dennis.com.cn:8082/gitlab/gitlab-ce:18.9.3-ce.0

    # 容器名称
    container_name: gitlab

    # Docker 重启后自动启动
    restart: always

    # 容器内部主机名
    hostname: gitlab.local

    environment:
      GITLAB_OMNIBUS_CONFIG: |
        ## 基础访问配置

        # GitLab Web访问地址
        external_url 'http://10.17.220.221:81'

        # Git SSH 克隆端口（宿主机映射2222）
        gitlab_rails['gitlab_shell_ssh_port'] = 2222

        ###############################################################
        ## 关闭非必要组件（降低资源占用）
        ###############################################################

        prometheus_monitoring['enable'] = false
        alertmanager['enable'] = false
        node_exporter['enable'] = false
        postgres_exporter['enable'] = false
        redis_exporter['enable'] = false
        gitlab_exporter['enable'] = false

        # Kubernetes Agent
        gitlab_kas['enable'] = false

        # Docker Registry
        registry['enable'] = false

        # GitLab Pages
        gitlab_pages['enable'] = false

        # 关闭邮件接收
        gitlab_rails['incoming_email_enabled'] = false

        ###############################################################
        ## 性能优化（适用于2C4G左右服务器）
        ###############################################################

        # Sidekiq 最大并发
        sidekiq['max_concurrency'] = 5

        # 禁止内存自动Kill
        sidekiq['memory_killer'] = false

        # Puma Worker
        puma['worker_processes'] = 1

        # PostgreSQL缓存
        postgresql['shared_buffers'] = '64MB'

    ###############################################################
    ## 端口映射
    ###############################################################
    ports:
      # Web访问
      - "81:81"

      # Git SSH Clone
      - "2222:22"

    ###############################################################
    ## 数据持久化
    ###############################################################
    volumes:
      # GitLab配置
      - /srv/gitlab/config:/etc/gitlab

      # GitLab日志
      - /srv/gitlab/logs:/var/log/gitlab

      # GitLab数据
      - /srv/gitlab/data:/var/opt/gitlab

    ###############################################################
    ## 共享内存
    ###############################################################
    shm_size: "512m"

    ###############################################################
    ## 网络
    ###############################################################
    networks:
      - gitlab

  #####################################################################
  # GitLab Runner
  gitlab-runner:
    # GitLab Runner 官方镜像
    image: nexus.dennis.com.cn:8082/gitlab/gitlab-runner:v18.9.0

    # 容器名称
    container_name: gitlab-runner

    # 自动重启
    restart: always

    # 等待 GitLab 启动
    depends_on:
      - gitlab

    ## 数据持久化
    volumes:
      # Runner配置文件
      # 注册Runner后自动生成config.toml
      - /srv/gitlab-runner/config:/etc/gitlab-runner

      # Docker Socket
      # Docker Executor必须挂载
      # CI/CD运行时调用宿主机Docker
      - /var/run/docker.sock:/var/run/docker.sock

    ## 环境变量
    environment:
      # Runner 默认使用 Docker Executor
      - DOCKER_HOST=unix:///var/run/docker.sock
    ## 网络
    networks:
      - gitlab
# Docker 网络
networks:
  gitlab:
    driver: bridge
```
## 分别启动
```
root@jasper:/home/jasper/gitlab# cat docker-compose.yaml 
version: '3'

services:
  gitlab:
    image: gitlab/gitlab-ce:19.1.0-ce.0
    container_name: gitlab
    restart: always
    hostname: 'gitlab.local'          # 容器内部主机名，可自定义
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        # ---------- 基础访问 ----------
        external_url 'http://10.17.220.221:81'
        gitlab_rails['gitlab_shell_ssh_port'] = 2222

        # ---------- 关闭所有非必需服务（极致精简） ----------
        # 监控与指标（Prometheus 全家桶）
        prometheus_monitoring['enable'] = false
        alertmanager['enable'] = false
        node_exporter['enable'] = false
        postgres_exporter['enable'] = false
        redis_exporter['enable'] = false
        gitlab_exporter['enable'] = false   # 新增，部分版本存在

        # Kubernetes Agent
        gitlab_kas['enable'] = false

        # Docker 镜像仓库
        registry['enable'] = false

        # GitLab Pages
        gitlab_pages['enable'] = false


        # 邮件接收（保留发送能力，仅关闭接收）
        gitlab_rails['incoming_email_enabled'] = false

        # ---------- 性能调优（低内存/CPU 环境） ----------
        sidekiq['max_concurrency'] = 5
        puma['worker_processes'] = 1
        postgresql['shared_buffers'] = '64MB'
        sidekiq['memory_killer'] = false      # 禁用 Sidekiq 内存杀手，避免意外重启


    ports:
      - '81:81'        # HTTP 访问
      - '2222:22'      # SSH（注意映射到宿主机的 2222 端口）
    volumes:
      - '/srv/gitlab/config:/etc/gitlab'   # 配置目录（请确保内部 gitlab.rb 为空或仅含必要项）
      - '/srv/gitlab/logs:/var/log/gitlab'
      - '/srv/gitlab/data:/var/opt/gitlab'
    shm_size: '512m'   # 共享内存，提升性能
    networks:
      - gitlab

networks:
  gitlab:
    driver: bridge
```
## 稍等5分钟
## runner
```
docker run -d --name gitlab-runner --restart always   -v /var/run/docker.sock:/var/run/docker.sock   -v /srv/gitlab-runner/config:/etc/gitlab-runner   gitlab/gitlab-runner:v19.1.0
```
## gitlab相关注意事项
### 重置密码（有点慢）
```
用户名 root
root@jasper:/home/jasper/gitlab# docker exec -it gitlab bash
root@gitlab:/# gitlab-rake "gitlab:password:reset[root]"
Enter password: 
Confirm password: 
Password successfully updated for user with username root.
root@gitlab:/# 
```
nexus.dennis.com.cn:8082/ubuntu:24.04
### 注册
```
root@jasper:/home/jasper/gitlab# docker exec -it gitlab-runner bash                         
                                                
root@c33bb25fada7:/# gitlab-runner register  --url http://10.17.220.221:81  --token glrt-gklllOazasAsc2DE6_433G86MQp0OjEKdToxCw.01.121qrhl2r
Runtime platform                                    arch=amd64 os=linux pid=37 revision=5eb085ab version=19.1.0
Running in system-mode.                            
                                                   
Enter the GitLab instance URL (for example, https://gitlab.com/):
[http://10.17.220.221:81]: http://10.17.220.221:81
Verifying runner... is valid                        correlation_id=01KVJ99VK54V5D2VE7FX0Z5JHD runner=gklllOaza runner_name=c33bb25fada7
Enter a name for the runner. This is stored only in the local config.toml file:
[c33bb25fada7]: Xieyuxuan
Enter an executor: ssh, shell, custom, docker, docker+machine, kubernetes, parallels, virtualbox, instance, docker-windows, docker-autoscaler:
docker
Enter the default Docker image (for example, ruby:3.3):
alpine:3.20
Runner registered successfully. Feel free to start it, but if it's running already the config should be automatically reloaded!
 
Configuration (with the authentication token) was saved in "/etc/gitlab-runner/config.toml
```