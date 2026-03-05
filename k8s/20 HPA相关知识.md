## Pod自动扩容
```
Pod自动扩容:可以根据cpu使用率或自定义指标，自动对pod进行扩缩容，前提是该对象必须配置了resources.requests.cpu或resources.requests.memory才可以
```
```
替换配置文件
kubectl replace -f nginx-deploy.yaml
```
```
执行命令：
kubectl autoscale deploy 名字 --cpu-percent=20 --min=2 --max=5
```
```
查看cpu和内存使用情况
kubectl top pod 名字（必须要有metrics）
kubectl top pod(查看所有)
```
```
安装metrics
git clone https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
kubectl get pods -n kube-system
可以下载后上传
```
```
修改yaml文件
containers.args --kubelet-insecure-tls
```
```
kubectl apply -f metrics-server-components.yaml
```
```
kubectl get po --A|grep metrics
```
```yaml
#测试（死循环）
#配置一个service
apiVersion:v1
kind: Service
metadata:
  name: nginx
  labels:
    app:nginx
spec:
  selector:
    app: nginx-deploy
  ports:
  - port； 80
    targetPort: 80
    name: web
type: NodePort
```
```
while true; do wget -q -O- http://svcip地址 > /dev/null ;done//在运行的node上面运行
```
```
重新查看
kubectl top po
kubectl get hpa
```
```
kubectl get endpoints(ep)
```
