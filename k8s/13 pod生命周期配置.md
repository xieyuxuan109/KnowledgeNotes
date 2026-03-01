## 配置示例
```yaml
terminationGracePeriodSeconds: 30 #当pod被删除时候给pod的容忍时间，默认30秒 container级别配置
lifecycle: #生命周期的配置
  postStart: #容器生命周期启动阶段做的事情，不一定在容器command之前运行
    exec:
      command:
      - sh
      - -c
      - "echo '<h1>pre stop</h1>' > /usr/share/nginx/html/prestop.html"
  preStop:
    command:
    - sh
    - -c
    - "sleep 50;echo 'sleep finished' >> /usr/share/nginx/html/prestop.html"
    
```