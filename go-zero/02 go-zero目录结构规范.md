## 多服务项目布局
```
mall/                         # 项目根目录
├── common/                   # 公共库（工具类）
└── service/                  # 服务存放目录
    ├── user/
    │   ├── api/              # user 的 HTTP 服务
    │   ├── rpc/              # user 的 RPC 服务
    │   └── model/            # user 专属的 model 层
    ├── order/
    │   ├── api/
    │   ├── rpc/
    │   └── model/            # order 专属的 model 层
    └── product/
        ├── api/
        ├── rpc/
        └── model/            # product 专属的 model 层
```