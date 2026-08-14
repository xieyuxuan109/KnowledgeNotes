## 进入容器查看config
```
root@dennis:/data/gitlab# docker ps|grep runner
d45787b1fded   gitlab/gitlab-runner:v18.9.0                      "/usr/bin/dumb-init …"   3 months ago   Up 8 hours                                                                                                           gitlab-gitlab-runner-1
root@dennis:/data/gitlab# docker exec -it d45 bash
root@d45787b1fded:/# cat /etc/gitlab-runner/
.runner_system_id  config.toml        
root@d45787b1fded:/# cat /etc/gitlab-runner/config.toml 
concurrent = 1
check_interval = 0
connection_max_age = "15m0s"
shutdown_timeout = 0

[session_server]
  session_timeout = 1800

[[runners]]
  name = "87261bdc4d54"
  url = "http://10.17.210.200:81"
  id = 2
  token = "glrtr-yUfbmpSsVmbT9qNBNp2R"
  token_obtained_at = 2026-04-05T12:24:38Z
  token_expires_at = 0001-01-01T00:00:00Z
  executor = "docker"
  [runners.cache]
    MaxUploadedArchiveSize = 0
    [runners.cache.s3]
    [runners.cache.gcs]
    [runners.cache.azure]
  [runners.docker]
    tls_verify = false
    image = "node:24.13"
    privileged = false
    disable_entrypoint_overwrite = false
    oom_kill_disable = false
    disable_cache = false
    volumes = ["/cache"]
    shm_size = 0
    network_mtu = 0

[[runners]]
  name = "golang"
  url = "http://10.17.210.200:81"
  id = 3
  token = "glrtr-bacPNFggDTcwFxCsgG_d"
  token_obtained_at = 2026-04-05T12:28:34Z
  token_expires_at = 0001-01-01T00:00:00Z
  executor = "docker"
  [runners.cache]
    MaxUploadedArchiveSize = 0
    [runners.cache.s3]
    [runners.cache.gcs]
    [runners.cache.azure]
  [runners.docker]
    tls_verify = false
    image = "golang:1.25"
    privileged = false
    disable_entrypoint_overwrite = false
    oom_kill_disable = false
    disable_cache = false
    volumes = ["/cache"]
    shm_size = 0
    network_mtu = 0
```
## 使用sed修改config
```
root@d45787b1fded:/# cd /etc/gitlab-runner/            
root@d45787b1fded:/etc/gitlab-runner# cp config.toml config.toml.bak
root@d45787b1fded:/etc/gitlab-runner# sed -i '/network_mtu/i\    pull_policy = "if-not-present"' config.toml          
```
## 验证
```
root@d45787b1fded:/etc/gitlab-runner# cat /etc/gitlab-runner/config.toml 
concurrent = 1
check_interval = 0
connection_max_age = "15m0s"
shutdown_timeout = 0

[session_server]
  session_timeout = 1800

[[runners]]
  name = "87261bdc4d54"
  url = "http://10.17.210.200:81"
  id = 2
  token = "glrtr-yUfbmpSsVmbT9qNBNp2R"
  token_obtained_at = 2026-04-05T12:24:38Z
  token_expires_at = 0001-01-01T00:00:00Z
  executor = "docker"
  [runners.cache]
    MaxUploadedArchiveSize = 0
    [runners.cache.s3]
    [runners.cache.gcs]
    [runners.cache.azure]
  [runners.docker]
    tls_verify = false
    image = "node:24.13"
    privileged = false
    disable_entrypoint_overwrite = false
    oom_kill_disable = false
    disable_cache = false
    volumes = ["/cache"]
    shm_size = 0
    pull_policy = "if-not-present"
    network_mtu = 0

[[runners]]
  name = "golang"
  url = "http://10.17.210.200:81"
  id = 3
  token = "glrtr-bacPNFggDTcwFxCsgG_d"
  token_obtained_at = 2026-04-05T12:28:34Z
  token_expires_at = 0001-01-01T00:00:00Z
  executor = "docker"
  [runners.cache]
    MaxUploadedArchiveSize = 0
    [runners.cache.s3]
    [runners.cache.gcs]
    [runners.cache.azure]
  [runners.docker]
    tls_verify = false
    image = "golang:1.25"
    privileged = false
    disable_entrypoint_overwrite = false
    oom_kill_disable = false
    disable_cache = false
    volumes = ["/cache"]
    shm_size = 0
    pull_policy = "if-not-present"
    network_mtu = 0
```