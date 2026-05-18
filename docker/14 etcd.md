### 配置
```
  etcd:
    image: quay.io/coreos/etcd:v3.5.17
    container_name: sec_etcd
    environment:
      ALLOW_NONE_AUTHENTICATION: yes
      ETCD_LISTEN_CLIENT_URLS: http://0.0.0.0:2379
      ETCD_ADVERTISE_CLIENT_URLS: http://10.17.220.221:2379
    ports:
      - "2379:2379"
      - "2380:2380"
```
### 测试
```
root@jasper:/home/jasper# curl http://127.0.0.1:2379/version
{"etcdserver":"3.5.17","etcdcluster":"3.5.0"}
```
### 服务注册后查看
```
root@jasper:/home/jasper# docker exec -it sec_etcd  etcdctl get --prefix order.rpc
order.rpc/7587894584484208389
100.64.164.2:8081
```
### 含义解释
```
order.rpc/7587894584484208389
100.64.164.2:8081
这表示你的 order.rpc 服务已经成功注册到了 etcd 中！

输出含义解释
Key：order.rpc/7587894584484208389
其中 order.rpc 是你配置的服务 Key，后面的数字是 etcd 自动生成的 租约 ID（或唯一标识），用于区分同一个服务的多个实例。

Value：100.64.164.2:8081
这是你的 RPC 服务实际监听的 IP 地址和端口（ListenOn: 0.0.0.0:8081 会绑定到容器内部的 IP 100.64.164.2）。其他服务通过 etcd 查询 order.rpc 就能得到这个地址，然后发起 gRPC 调用。
```