## docker里面再运行docker
```
docker run --privileged --name dind-docker -d docker:stable-dind
```
## 进入容器 注意只能是sh
```
docker exec -it 1bb sh
```
## gitlab运行的话需要修改配置
```
GitLab Runner 开启 privileged
root@dennis:~# docker exec -it e1c bash
root@e1cc5df74e40:/etc/gitlab-runner# cat /etc/gitlab-runner/config.toml|grep privileged
    privileged = false
    privileged = false
root@e1cc5df74e40:/etc/gitlab-runner# sed -i 's/privileged = false/privileged = true/g' config.toml
root@e1cc5df74e40:/etc/gitlab-runner# cat /etc/gitlab-runner/config.toml|grep privileged
    privileged = true
    privileged = true
```
## 继续修改配置
```
root@528e5c3b3bef:/# cat /etc/gitlab-runner/config.toml|grep volumes
    volumes = ["/cache"]
    volumes = ["/cache"]
root@528e5c3b3bef:/# cd /etc/gitlab-runner/           
root@528e5c3b3bef:/etc/gitlab-runner# sed -i 's/volumes = ["/cache"]/volumes = ["/var/run/docker.sock:/var/run/docker.sock", "/etc/docker/daemon.json:/etc/docker/daemon.json", "/cache"]/g' config.toml
sed: -e expression #1, char 37: unknown option to `s'
root@528e5c3b3bef:/etc/gitlab-runner# sed -i 's/volumes = ["^Cache"]/volumes = ["/var/run/docker.sock:/var/run/docker.sock", "/etc/docker/daemon.json:/etc/docker/daemon.json", "/cache"]/g' config.toml
root@528e5c3b3bef:/etc/gitlab-runner# sed -i 's|volumes = \["/cache"\]|volumes = ["/var/run/docker.sock:/var/run/docker.sock", "/etc/docker/daemon.json:/etc/docker/daemon.json", "/cache"]|g' config.toml   
root@528e5c3b3bef:/etc/gitlab-runner# cat /etc/gitlab-runner/config.toml|grep volumes
    volumes = ["/var/run/docker.sock:/var/run/docker.sock", "/etc/docker/daemon.json:/etc/docker/daemon.json", "/var/lib/docker/:/var/lib/docker/","/cache"]
    volumes = ["/var/run/docker.sock:/var/run/docker.sock", "/etc/docker/daemon.json:/etc/docker/daemon.json","/var/lib/docker/:/var/lib/docker/","/cache"]
```
## 继续修改
```
sed -i 's|"/cache"|"/cache","root/.kube:root/.kube"|g' config.toml
```
## 替换
```
sed -i 's|"root/.kube:root/.kube"|"/root/.kube:/root/.kube"|g' config.toml
sed -i 's|"/root/.kube:/root/.kube"|"/root/.kube:/tmp/.kube"|g' config.toml
sed -i 's|"/root/.kube:/tmp/.kube"|"/root/.kube:/tmp/.kube","/usr/bin/docker:/usr/bin/docker"|g'
```