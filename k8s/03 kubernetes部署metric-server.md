### metric-server 仓库地址
```
https://github.com/kubernetes-sigs/metrics-server/releases
```
### 下载yaml文件
```
https://release-assets.githubusercontent.com/github-production-release-asset/92132038/0626adef-e098-4155-ab3f-6f67afd3bce4?sp=r&sv=2018-11-09&sr=b&spr=https&se=2025-12-15T03%3A15%3A16Z&rscd=attachment%3B+filename%3Dcomponents.yaml&rsct=application%2Foctet-stream&skoid=96c2d410-5711-43a1-aedd-ab1947aa7ab0&sktid=398a6654-997b-47e9-b12b-9515b896b4de&skt=2025-12-15T02%3A15%3A14Z&ske=2025-12-15T03%3A15%3A16Z&sks=b&skv=2018-11-09&sig=k1f%2FYHGbWb4pNTVjamzuo0gUzQb68%2F2MlJ%2FZKvQdm40%3D&jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmVsZWFzZS1hc3NldHMuZ2l0aHVidXNlcmNvbnRlbnQuY29tIiwia2V5Ijoia2V5MSIsImV4cCI6MTc2NTc2NTk4MiwibmJmIjoxNzY1NzY1NjgyLCJwYXRoIjoicmVsZWFzZWFzc2V0cHJvZHVjdGlvbi5ibG9iLmNvcmUud2luZG93cy5uZXQifQ.VG38xsx2KzuhX2PDwYX82u6KwPk88r_Teq-Fn3plsMM&response-content-disposition=attachment%3B%20filename%3Dcomponents.yaml&response-content-type=application%2Foctet-stream
```
### 修改metric-server.yaml镜像名称
```
...
image: registry.cn-hangzhou.aliyuncs.com/google_containers/metrics-server:v0.8.0
...
```
### args 中添加 - --kubelet-insecure-tls 禁用证书验证
```
...
      - args:
        - --cert-dir=/tmp
        - --secure-port=10250
        - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
        - --kubelet-use-node-status-port
        - --metric-resolution=15s
        - --kubelet-insecure-tls
...
```
### 部署
```
# kubectl apply -f metric-server.yaml 
serviceaccount/metrics-server unchanged
clusterrole.rbac.authorization.k8s.io/system:aggregated-metrics-reader unchanged
clusterrole.rbac.authorization.k8s.io/system:metrics-server unchanged
rolebinding.rbac.authorization.k8s.io/metrics-server-auth-reader unchanged
clusterrolebinding.rbac.authorization.k8s.io/metrics-server:system:auth-delegator unchanged
clusterrolebinding.rbac.authorization.k8s.io/system:metrics-server unchanged
service/metrics-server unchanged
deployment.apps/metrics-server created
apiservice.apiregistration.k8s.io/v1beta1.metrics.k8s.io unchanged
```
### 查看pod运行状态
```
# kubectl get pods -A|grep metric
kube-system                metrics-server-6678d74bdd-t4gh6            1/1     Running   0             16m
```
### 测试
```
# kubectl top node
NAME     CPU(cores)   CPU(%)   MEMORY(bytes)   MEMORY(%)   
master   191m         11%      2134Mi          67%         
node1    102m         6%       1336Mi          41%         
node2    83m          5%       956Mi           30%         
# kubectl top pods -A
NAMESPACE                  NAME                                       CPU(cores)   MEMORY(bytes)   
default                    busybox-with-probes-9dd477544-b9w2l        1m           0Mi             
default                    busybox-with-probes-9dd477544-xhqj6        1m           3Mi             
default                    test-nginx-547d6dfb9c-2fwp7                0m           3Mi             
default                    test-nginx-547d6dfb9c-74bqp                0m           3Mi             
default                    test-nginx-547d6dfb9c-bnffr                0m           3Mi             
default                    test-nginx-547d6dfb9c-h26zl                0m           10Mi            
default                    test-nginx-547d6dfb9c-nq546                0m           3Mi             
default                    test-nginx-547d6dfb9c-zgr2h                0m           3Mi             
default                    test-nginx1-69d6d549d7-2wsdj               0m           6Mi             
default                    test-nginx1-69d6d549d7-hs7b9               0m           6Mi             
extension-metrics-server   metrics-server-frontend-d4b5c8c9b-cn2t6    1m           3Mi             
kube-system                calico-kube-controllers-678fc69664-l48p5   3m           24Mi            
kube-system                calico-node-bkflp                          41m          136Mi           
kube-system                calico-node-k8gkg                          54m          131Mi           
kube-system                calico-node-zkqd4                          41m          147Mi           
kube-system                coredns-7ddb55578d-hc6h8                   2m           22Mi            
kube-system                coredns-7ddb55578d-hxfbk                   3m           19Mi            
kube-system                kube-apiserver-master                      55m          376Mi           
kube-system                kube-controller-manager-master             23m          79Mi            
kube-system                kube-proxy-kvnfs                           1m           25Mi            
kube-system                kube-proxy-kzbdl                           17m          21Mi            
kube-system                kube-proxy-qfshx                           8m           33Mi            
kube-system                kube-scheduler-master                      11m          35Mi            
kube-system                metrics-server-6678d74bdd-t4gh6            5m           18Mi            
kube-system                nodelocaldns-g2jgt                         3m           19Mi            
kube-system                nodelocaldns-mtcdd                         6m           13Mi            
kube-system                nodelocaldns-wjmvf                         6m           13Mi            
kubesphere-system          extensions-museum-f44c598d8-d8g9n          1m           26Mi            
kubesphere-system          ks-apiserver-c4d78ddf4-hxmfl               1m           41Mi            
kubesphere-system          ks-console-6bfbc957cb-w8v4k                1m           77Mi            
kubesphere-system          ks-console-embed-5597c7df68-m22rk          1m           3Mi             
kubesphere-system          ks-controller-manager-784cdb8c7-bbjnz      9m           88Mi            
kubesphere-system          ks-posthog-dc697467d-mxhcs                 0m           15Mi 
```