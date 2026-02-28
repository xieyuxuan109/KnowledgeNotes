# Kubernetes Pod 三种探针说明

K8s Pod 有三种健康探针（Probe）：

- Liveness Probe（存活探针）
- Readiness Probe（就绪探针）
- Startup Probe（启动探针）

---

## Liveness Probe（存活探针）

**作用：** 判断容器是否“死掉”。

如果探测失败：

- kubelet 会重启容器

### 使用场景

- 程序死锁
- 线程卡死
- 内存异常
- 服务卡住但进程仍在

### 示例

```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: 8080
  initialDelaySeconds: 30
  periodSeconds: 10
  failureThreshold: 3
```

如果连续失败达到阈值：

 容器会被 kill 并自动重启

---

##  Readiness Probe（就绪探针）

**作用：** 判断容器是否可以接收流量。

如果探测失败：

- 不会重启容器
- 会从 Service 负载均衡列表中移除

### 使用场景

- 应用正在启动
- 正在加载缓存
- 依赖的数据库未连接成功
- 应用正在优雅下线

### 示例

```yaml
readinessProbe:
  httpGet:
    path: /ready
    port: 8080
  periodSeconds: 5
  failureThreshold: 3
```

如果失败：

 Service 不再转发流量到该 Pod  
 Pod 仍然运行

---

##  Startup Probe（启动探针）

**作用：** 解决应用启动慢的问题。

在 Startup Probe 成功之前：

- Liveness 和 Readiness 不会生效

### 使用场景

- Java 应用启动慢
- SpringBoot 冷启动慢
- 大型初始化加载

### 示例

```yaml
startupProbe:
  httpGet:
    path: /healthz
    port: 8080
  failureThreshold: 30
  periodSeconds: 10
```

允许最大启动时间：

30 × 10 = 300 秒

---

## 三种探针对比

| 探针 | 失败后行为 | 是否重启容器 | 用途 |
|------|------------|--------------|------|
| Liveness | 重启容器 | ✅ | 解决程序假死 |
| Readiness | 不接收流量 | ❌ | 控制流量 |
| Startup | 阻止前两者生效 | ❌ | 解决慢启动 |

---

## 通俗理解

- Liveness = 你还活着吗？
- Readiness = 你现在能接流量吗？
- Startup = 你启动完成了吗？

---

## 推荐实践

- 普通 Web 服务：配置 Liveness + Readiness
- 慢启动应用：加 Startup Probe
- 只想控制流量：只使用 Readiness