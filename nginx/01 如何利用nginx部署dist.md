## 迁移dist文件
```
root@jasper:/home/jasper/blogx_server/init/deploy/blogx_web# docker cp ./dist/ deploy-blogx_web-1:/usr/share/nginx/html
Successfully copied 8.4MB to deploy-blogx_web-1:/usr/share/nginx/html
```
```
cat > /usr/share/nginx/html/index.html << EOF
<h1>Hello Debug</h1>
<p>This is temporary</p>
EOF
```