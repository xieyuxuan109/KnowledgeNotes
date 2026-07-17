## 克隆 Dify
```
root@jasper:/home/jasper# git clone --branch "$(curl -s https://api.github.com/repos/langgenius/dify/releases/latest | jq -r .tag_name)" https://github.com/langgenius/dify.git
Cloning into 'dify'...
remote: Enumerating objects: 460178, done.
remote: Counting objects: 100% (510/510), done.
remote: Compressing objects: 100% (313/313), done.
remote: Total 460178 (delta 315), reused 285 (delta 192), pack-reused 459668 (from 3)
Receiving objects: 100% (460178/460178), 401.83 MiB | 6.68 MiB/s, done.
Resolving deltas: 100% (336719/336719), done.
Note: switching to '3aa26fb6374bbd47e5469f7d7cc25f3e0075a60c'.

You are in 'detached HEAD' state. You can look around, make experimental
changes and commit them, and you can discard any commits you make in this
state without impacting any branches by switching back to a branch.

If you want to create a new branch to retain commits you create, you may
do so (now or later) by using -c with the switch command. Example:

  git switch -c <new-branch-name>

Or undo this operation with:

  git switch -

Turn off this advice by setting config variable advice.detachedHead to false

Updating files: 100% (13013/13013), done.
```
## 启动 Dify
```
root@jasper:/home/jasper# cd dify/docker
root@jasper:/home/jasper/dify/docker# cp .env.example .env
root@jasper:/home/jasper/dify/docker# docker compose up -d
[+] up 15/15
 ✔ Network docker_default              Created                                                                                                                                                                             0.1ss
 ✔ Network docker_ssrf_proxy_network   Created                                                                                                                                                                             0.1ss
 ✔ Container docker-redis-1            Created                                                                                                                                                                             0.5ss
 ✔ Container docker-init_permissions-1 Exited                                                                                                                                                                              9.1ss
 ✔ Container docker-db_postgres-1      Healthy                                                                                                                                                                             11.9s
 ✔ Container docker-web-1              Created                                                                                                                                                                             0.5ss
 ✔ Container docker-weaviate-1         Created                                                                                                                                                                             0.4ss
 ✔ Container docker-sandbox-1          Created                                                                                                                                                                             0.5ss
 ✔ Container docker-ssrf_proxy-1       Created                                                                                                                                                                             0.4ss
 ✔ Container docker-api-1              Created                                                                                                                                                                             0.3ss
 ✔ Container docker-api_websocket-1    Created                                                                                                                                                                             0.3ss
 ✔ Container docker-worker-1           Created                                                                                                                                                                             0.3ss
 ✔ Container docker-worker_beat-1      Created                                                                                                                                                                             0.3ss
 ✔ Container docker-plugin_daemon-1    Created                                                                                                                                                                             0.2ss
 ✔ Container docker-nginx-1            Created                                                                                                                                                                             0.2ss
```

## 检验
```
docker compose ps
NAME                     IMAGE                                       COMMAND                  SERVICE         CREATED          STATUS                             PORTS
docker-api-1             langgenius/dify-api:1.10.1                  "/bin/bash /entrypoi…"   api             26 seconds ago   Up 22 seconds                      5001/tcp
docker-db_postgres-1     postgres:15-alpine                          "docker-entrypoint.s…"   db_postgres     26 seconds ago   Up 25 seconds (healthy)            5432/tcp
docker-nginx-1           nginx:latest                                "sh -c 'cp /docker-e…"   nginx           26 seconds ago   Up 22 seconds                      0.0.0.0:80->80/tcp, :::80->80/tcp, 0.0.0.0:443->443/tcp, :::443->443/tcp
docker-plugin_daemon-1   langgenius/dify-plugin-daemon:0.4.1-local   "/bin/bash -c /app/e…"   plugin_daemon   26 seconds ago   Up 22 seconds                      0.0.0.0:5003->5003/tcp, :::5003->5003/tcp
docker-redis-1           redis:6-alpine                              "docker-entrypoint.s…"   redis           26 seconds ago   Up 25 seconds (health: starting)   6379/tcp
docker-sandbox-1         langgenius/dify-sandbox:0.2.12              "/main"                  sandbox         26 seconds ago   Up 25 seconds (health: starting)   
docker-ssrf_proxy-1      ubuntu/squid:latest                         "sh -c 'cp /docker-e…"   ssrf_proxy      26 seconds ago   Up 25 seconds                      3128/tcp
docker-weaviate-1        semitechnologies/weaviate:1.27.0            "/bin/weaviate --hos…"   weaviate        26 seconds ago   Up 25 seconds                      
docker-web-1             langgenius/dify-web:1.10.1                  "/bin/sh ./entrypoin…"   web             26 seconds ago   Up 25 seconds                      3000/tcp
docker-worker-1          langgenius/dify-api:1.10.1                  "/bin/bash /entrypoi…"   worker          26 seconds ago   Up 22 seconds                      5001/tcp
docker-worker_beat-1     langgenius/dify-api:1.10.1                  "/bin/bash /entrypoi…"   worker_beat     26 seconds ago   Up 22 seconds                      5001/tcp
```