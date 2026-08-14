## docker save镜像
```
docker save -o myimage.tar docker:stable-dind
```
## 传递镜像
```
root@jasper:~# ls
myimage.tar  snap
root@jasper:~# scp myimage.tar root@10.17.210.200:/root
The authenticity of host '10.17.210.200 (10.17.210.200)' can't be established.
ED25519 key fingerprint is SHA256:lA33u9A0PRA7XV6qESwwQefQlNpsct+oy8gYScdurdM.
This key is not known by any other names.
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added '10.17.210.200' (ED25519) to the list of known hosts.
root@10.17.210.200's password: 
Permission denied, please try again.
root@10.17.210.200's password: 
myimage.tar                                                                                                                                                                                   100%   70MB  11.0MB/s   00:06    
```
## 解压镜像
```
root@dennis:~# docker load -i myimage.tar
Loaded image: docker:stable-dind
```
```
cat > .dockerignore <<EOF
node_modules
dist
.git
EOF
```