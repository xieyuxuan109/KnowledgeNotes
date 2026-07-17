## 提交代码
```
git config --global user.name "你的名字"
git config --global user.email "你的邮箱@example.com"
```
## 生成密钥
```
root@jasper:/home/jasper/gitlab# ssh-keygen -t ed25519 -C "xieyuxuan109@163.com"
Generating public/private ed25519 key pair.
Enter file in which to save the key (/root/.ssh/id_ed25519): /root/.ssh/id_ed25519
Enter passphrase (empty for no passphrase): 
Enter same passphrase again: 
Your identification has been saved in /root/.ssh/id_ed25519
Your public key has been saved in /root/.ssh/id_ed25519.pub
The key fingerprint is:
SHA256:oujLS3F20Xx6Xo0cjZW3W6OMN4MRKFvFZuo9/ENljIg xieyuxuan109@163.com
The key's randomart image is:
+--[ED25519 256]--+
|          +.  .. |
|       + o = +. .|
|      . * * = +..|
|       o E + = *.|
|  . o o S + O = +|
|   = o . + B B . |
|  o .     . = o  |
| +           o   |
|  =o          .  |
+----[SHA256]-----+
root@jasper:/home/jasper/gitlab# cat /root/.ssh/id_ed25519 //私钥
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAzc2gtZW
QyNTUxOQAAACANgQ7jPhJofBpQ6fii9+4PLSuK0PHYEsWe+cVyDfB/DAAAAJiRU1jskVNY
7AAAAAtzc2gtZWQyNTUxOQAAACANgQ7jPhJofBpQ6fii9+4PLSuK0PHYEsWe+cVyDfB/DA
AAAEDbGisogAJNi+WojJumDLJ3Q+O0+LGps6NLHIIPvvVuaA2BDuM+Emh8GlDp+KL37g8t
K4rQ8dgSxZ75xXIN8H8MAAAAFHhpZXl1eHVhbjEwOUAxNjMuY29tAQ==
-----END OPENSSH PRIVATE KEY-----
root@jasper:/home/jasper/gitlab# cat /root/.ssh/id_ed25519.pub //公钥 粘贴这个
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIA2BDuM+Emh8GlDp+KL37g8tK4rQ8dgSxZ75xXIN8H8M xieyuxuan109@163.com
```
## 添加代码仓库
```
root@jasper:/home/jasper# git clone ssh://git@10.17.220.221:2222/xieyuxuan/blogx.git
Cloning into 'blogx'...
The authenticity of host '[10.17.220.221]:2222 ([10.17.220.221]:2222)' can't be established.
ED25519 key fingerprint is SHA256:Sq+joHFyEdnB2+qSE67U4MCEr8CG6RY2gpqEKP7+bi0.
This key is not known by any other names.
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added '[10.17.220.221]:2222' (ED25519) to the list of known hosts.
warning: You appear to have cloned an empty repository.
```
```
root@jasper:/home/jasper/use-gitlab-runner-as-cd-agent# git checkout -b test
Switched to a new branch 'test'
root@jasper:/home/jasper/use-gitlab-runner-as-cd-agent# git push
fatal: The current branch test has no upstream branch.
To push the current branch and set the remote as upstream, use

    git push --set-upstream origin test

To have this happen automatically for branches without a tracking
upstream, see 'push.autoSetupRemote' in 'git help config'.

root@jasper:/home/jasper/use-gitlab-runner-as-cd-agent# git push --set-upstream origin test
Username for 'http://10.17.220.221:81': root
Password for 'http://root@10.17.220.221:81': 
Enumerating objects: 31, done.
Counting objects: 100% (31/31), done.
Delta compression using up to 5 threads
Compressing objects: 100% (21/21), done.
Writing objects: 100% (31/31), 61.50 KiB | 2.93 MiB/s, done.
Total 31 (delta 0), reused 0 (delta 0), pack-reused 0
remote: 
remote: To create a merge request for test, visit:
remote:   http://10.17.220.221:81/develop/test/-/merge_requests/new?merge_request%5Bsource_branch%5D=test
remote: 
To http://10.17.220.221:81/develop/test.git
 * [new branch]      test -> test
branch 'test' set up to track 'origin/test'.
```