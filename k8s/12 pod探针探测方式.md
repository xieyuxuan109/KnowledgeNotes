# Kubernetes Probe 探测方式

Kubernetes 的三种 Probe（Liveness / Readiness / Startup）支持三种探测方式：

- HTTP
- TCP
- Exec

---

## 1️⃣ HTTP 探测（httpGet）

通过 HTTP 请求指定路径，根据返回状态码判断健康状态。

- 2xx / 3xx → 成功
- 其他状态码 → 失败

### 示例

```yaml
livenessProbe:
  httpGet:
    path: /healthz
    port: 8080
    scheme: HTTP
  initialDelaySeconds: 10
  periodSeconds: 10
```

### 常用字段

| 字段 | 说明 |
|------|------|
| path | 健康检查路径 |
| port | 端口 |
| scheme | HTTP / HTTPS |
| httpHeaders | 自定义请求头 |

---

## 2️⃣ TCP 探测（tcpSocket）

检测指定端口是否可以建立 TCP 连接。

只要端口能连通，就认为成功。

### 示例

```yaml
readinessProbe:
  tcpSocket:
    port: 3306
  periodSeconds: 5
```

### 适用场景

- MySQL
- Redis
- 不提供 HTTP 健康接口的服务

---

## 3️⃣ Exec 探测（exec）

在容器内部执行命令，根据返回码判断：

- 返回码 0 → 成功
- 非 0 → 失败

### 示例

```yaml
livenessProbe:
  exec:
    command:
      - cat
      - /tmp/healthy
  periodSeconds: 5
```

或者：

```yaml
readinessProbe:
  exec:
    command:
      - sh
      - -c
      - "pgrep java"
```

### 适用场景

- 需要自定义复杂逻辑
- 检查进程是否存在
- 检查文件或脚本结果

---

# 通用参数说明

| 参数 | 说明 |
|------|------|
| initialDelaySeconds | 容器启动后延迟多久开始探测 |
| periodSeconds | 探测间隔时间 |
| timeoutSeconds | 单次探测超时时间 |
| failureThreshold | 连续失败多少次判定失败 |
| successThreshold | 连续成功多少次判定成功（readiness 常用） |

---

# 三种方式对比

| 方式 | 优点 | 缺点 | 推荐程度 |
|------|------|------|----------|
| HTTP | 最常用、语义清晰 | 需要应用支持接口 | ⭐⭐⭐⭐⭐ |
| TCP | 简单 | 只能判断端口是否打开 | ⭐⭐⭐ |
| Exec | 灵活 | 消耗资源稍高 | ⭐⭐⭐⭐ |

---

# 推荐实践

- Web 服务 → 优先使用 HTTP
- 数据库 / 缓存 → 使用 TCP
- 特殊业务逻辑 → 使用 Exec

# 常见参数配置
```
initialDelaySeconds: 60 #初始化时间
timeoutSeconds: 2 #超时时间
successThreshold: 1 #检测1次成功就表示成功
failureThreshold: 2 # 检测失败两次就表示失败
```