## 创建
```
//generic 通用密码 一定要加单引号 避免特殊字符不能正确加密和解码
kubectl create secret generic my-secret --from-literal=username='admin' --from-literal=password=xieyuxuan
//查看
kubectl describe secret my-secret
```
## 基于base64的编码和解码
```
root@dennis:/home/dennis# echo 'hello'|base64
aGVsbG8K
root@dennis:/home/dennis# echo 'aGVsbG8K'|base64  --decode
hello
```
## 特殊字符
```
root@dennis:/home/dennis# echo 'hello@!-'|base64
aGVsbG9AIS0K
root@dennis:/home/dennis# echo 'aGVsbG9AIS0K'|base64 --decode
hello@!-
root@dennis:/home/dennis# echo hello@!-|base64
bash: !-: event not found
```
## 也可以通过文件创建
```
//查看详细命令
kubectl create secret -h
```