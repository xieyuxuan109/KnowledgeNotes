## 项目位置
```
root@dennis:~# docker exec -it a36 bash
root@runner-bacpnfggd-project-3-concurrent-0:/# ls
bin  boot  builds  cache  dev  etc  home  lib  lib64  media  mnt  opt  proc  root  run	sbin  srv  sys	tmp  usr  var
root@runner-bacpnfggd-project-3-concurrent-0:/# find / -name "VueCom"
root@runner-bacpnfggd-project-3-concurrent-0:/# ls builds/           
develop
root@runner-bacpnfggd-project-3-concurrent-0:/# ls builds/develop/
testhttp  testhttp.tmp
root@runner-bacpnfggd-project-3-concurrent-0:/# ls builds/develop/testhttp
README.md  env.d.ts  index.html  index.js  package-lock.json  package.json  public  src  tsconfig.app.json  tsconfig.json  tsconfig.node.json  vite.config.ts
```