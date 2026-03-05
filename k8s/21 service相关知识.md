## 标准
```
apiVersion: v1
kind: Service #资源类型为Service
metadata: 
  name: nginx-svc # Service名字
  labels:
    app: nginx # Service自己本身的标签
spec:
  selector: #匹配哪些pod会被该service代理
    app: nginx-deploy #所有匹配到这些标签的pod都可以通过该service进行访问
  ports: #端口映射
  - port: 80 #service自己的端口，子啊使用内网ip时候访问使用
    targetPort: 80 #目标pod的端口
    name: web #为端口起一个一个自己的名字，映射到ports中的端口，该端口是直接绑定在node上的，且集群中的每一个node都会绑定这个端口
    #也可以用于将服务暴露给外部访问，但是这种方式实际生产环境不推荐，效率较低。而且service是四层负载
```
```
查看
kubectl get svc nginx-svc
```
```
测试
kubectl exec -it busybox -- sh
wget http://nginx-svc(.default命名空间，跨命名空间访问)
```
```yaml
#标准endpoint yaml
apiVersion: v1        # API版本，Endpoints 属于核心API组 v1
kind: Endpoints       # 资源类型，表示服务后端地址列表

metadata:
  name: nginx-svc     # Endpoints 名称，必须和对应的 Service 名字相同
  namespace: default  # 所属命名空间，默认是 default

subsets:              # endpoint集合，可以包含多个地址组
- addresses:          # 后端服务地址列表（通常是 Pod IP）
  - ip: 10.244.1.3    # 第一个后端 Pod 的 IP 地址
  - ip: 10.244.2.5    # 第二个后端 Pod 的 IP 地址

  ports:              # 后端服务端口列表
  - port: 80          # Pod 提供服务的端口
    protocol: TCP     # 使用的协议，默认 TCP
```