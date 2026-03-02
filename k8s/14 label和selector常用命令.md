## 常用命令
```
下面以pod举例，其他组件也可以参考
```
### 如何添加label
```
永久添加label
修改组件的yaml配置文件，在metadata里面添加label
临时添加label
kubectl label po <资源名称> app=hello (-n 命名空间)
如何临时修改label
kubectl label po <资源名称> app=hello2 (-n 命名空间) --overwrite
```
### 查看和匹配label
```
查看
kubectl get po (-n 命名空间) --show-labels(查看labels)
匹配
kubectl get po (-A在所有命名空间里面匹配) -l app=hello(匹配的label)
多重匹配(&&的关系)
kubectl get po -l 'test!=1.0.1,type=app,author in (xxx,xxx,xiaoliu,xieyuxuan)'
```