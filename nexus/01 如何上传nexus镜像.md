docker login nexus.dennis.com.cn:8083
docker tag  nexus.dennis.com.cn:8083/<repository>/<image-name>:<tag>
docker push <nexus-server-url>/<repository>/<image-name>:<tag>
job_name:
  script:
    
echo 'admin' | docker login nexus.dennis.com.cn:8083
sudo hexo g 
spawn sudo -S hexo d
expect "password:"
send "123456"