# RBAC

- [概述](#overview)
- [原理](#principle)
- [设计](#design)
    - [数据库设计](#design-db)
    - [端点设计](#design-peer)
    - [rpc设计](#design-rpc)
    - [中间件设计](#design-middleware)
- [函数索引](#func-index)
- [客户端sdk](#client-sdk)

## <a id="overview">概述</a>

rbac(基于角色的访问控制)用于解决服务的访问控制问题. 例如: 一个公司的saas后台要求不同的员工对内部的服务有不同的访问控制和操作权限, 同时能够灵活控制. rbac 通过拆分出: 用户 -> 角色 -> 操作权限 -> 资源 来灵活控制用户对资源的访问. 对于一个常规的系统而言, 通常都有账户管理服务, 因此这里记录并实现一个简单易用的 rbac 服务.

## <a id="principle">原理</a>



## <a id="design">设计</a>

**目录结构**

```bash
./rbac
├── model
│   ├── dao.go <- 数据访问结构
│   ├── dpo.go <- 数据处理结构
│   └── dto.go <- 数据传输对象
├── repo.go <- 数据库交互
├── rest
│   ├── client.go <- restful 客户端 sdk
│   ├── handler.go <- 端点处理器
│   └── router.go <- 路由器
├── rpc
│   ├── client.go <- rpc 客户端 sdk
│   └── server.go <- rpc 服务器
└── service.go <- 业务逻辑
```

### <a id="design-db">数据库设计</a>

*table-name-1*

|字段|类型|null mode|默认值|约束|注释|
|:-:|:-:|:-:|:-:|:-:|:-:|
| - | - | - | - | - | - |

---

*table-name-2*

|字段|类型|null mode|默认值|约束|注释|
|:-:|:-:|:-:|:-:|:-:|:-:|
| - | - | - | - | - | - |

### <a id="design-peer">端点设计</a>

*title-1*

```
<method> <uri> <scheme>

<headers>...

<body>
```

|参数|参数类型|类型|约束|默认值|描述|
|:-:|:-:|:-:|:-:|:-:|:-:|
| - | - | - | - | - | - |

---

*title-2*

```
<method> <uri> <scheme>

<headers>...

<body>
```

|参数|参数类型|类型|约束|默认值|描述|
|:-:|:-:|:-:|:-:|:-:|:-:|
| - | - | - | - | - | - |

### <a id="design-rpc">rpc设计</a>



### <a id="design-middleware">中间件设计</a>

## <a id="func-index">函数索引</a>

## <a id="client-sdk">客户端sdk</a>