# RBAC

- [概述](#overview)
- [原理](#principle)
- [设计](#design)
    - [数据库设计](#design-db)
    - [端点设计](#design-peer)
    - [rpc设计](#design-rpc)
    - [中间件设计](#design-middleware)
    - [配置设计](#design-config)
- [函数索引](#func-index)
- [客户端sdk](#client-sdk)

## <a id="overview">概述</a>

rbac(基于角色的访问控制)用于解决服务的访问控制问题. 例如: 一个公司的saas后台要求不同的员工对内部的服务有不同的访问控制和操作权限, 同时能够灵活控制. rbac 通过拆分出: 用户 -> 角色 -> 操作权限 -> 资源 来灵活控制用户对资源的访问. 对于一个常规的系统而言, 通常都有账户管理服务, 因此这里记录并实现一个简单易用的 rbac 服务.

## <a id="principle">原理</a>

RBAC 即基于角色的权限访问控制（Role-Based Access Control). 这是一种通过角色关联权限，角色同时又关联用户的授权方式.

**基本概念**

- 角色（Role）：角色是一组权限的集合，代表了在组织内执行特定任务或职责的用户群体. 例如，“管理员”、“编辑”和“访客”等.
- 用户（User）：系统中的实际用户，他们被分配到一个或多个角色，从而获得相应的权限.
- 权限（Permission）：权限定义了用户可以对系统资源执行的操作，如“读取”、“写入”或“删除”等.
- 会话（Session）：用户通过身份验证后建立的与系统之间的连接，会话中用户的角色和权限将被激活.

**3 种模型**

rbac0: 最简单、最原始的实现方式.

rbac1: 在 rbac0 基础上加入了角色权限继承的功能.

rbac2: 在 rbac0 基础上加入了角色互斥制约的功能.

**思路**

1. 将 permission 细化为: action + resource
2. 集成 rbac1 思路: 每个 role 有 parent_id 形成继承追溯链路
3. 简单链路: 用户注册(uid, password, uname), 当前无任何角色 -> 赋予角色(关联permissions) -> 用户登录(会话内加载权限, 前端加载查看权限约束ui显示, 后端加载其他权限) -> 用户操作 -> 访问控制中间件 -> 业务逻辑.

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

**表设计**

*meta* - 元信息(所有的表都包含的字段, 统一提取并定义)

|字段|类型|null mode|默认值|约束|注释|
|:-:|:-:|:-:|:-:|:-:|:-:|
| id | bigint | not null | - | auto_inc, pk | 主键 |
| created_at | datetime | not null | current_timestamp() | - | 创建时间 |
| updated_at | datetime | not null | current_timestamp() | on update current_timestamp() | 更新时间 |
| deleted_at | datetime | null | null | - | 删除时间 |

---

*users* - 用户

|字段|类型|null mode|默认值|约束|注释|
|:-:|:-:|:-:|:-:|:-:|:-:|
| uid | bigint | not null | - | - | 用户唯一标识 |
| uname | varchar(32) | not null | - | - | 用户名 |
| password | varchar(255) | not null | - | - | 密码 |
| email | varchar(128) | not null | "" | - | 邮箱 |
| phone | varchar(128) | not null | "" | - | 手机 |

---

*roles* - 角色

|字段|类型|null mode|默认值|约束|注释|
|:-:|:-:|:-:|:-:|:-:|:-:|
| rid | bigint | not null | - | - | 角色唯一标识 |
| rname | varchar(255) | not null | - | - | 角色名称 |
| creater | bigint | not null | - | ref users.uid | 创建者id |
| inherit_id | bigint | not null | 0 | ref roles.rid | 继承id, 0-继承终点 |

---

*users_to_roles* - 用户-角色关联

|字段|类型|null mode|默认值|约束|注释|
|:-:|:-:|:-:|:-:|:-:|:-:|
| user_id | bigint | not null | - | ref users.uid | 用户id |
| role_id | bigint | not null | - | ref roles.rid | 角色id |
| creater | bigint | not null | - | ref users.uid | 创建者id |

---

*actions* - 操作

|字段|类型|null mode|默认值|约束|注释|
|:-:|:-:|:-:|:-:|:-:|:-:|
| aid | bigint | not null | - | - | 操作唯一标识 |
| aname | varchar(64) | not null | - | - | 操作名称 |
| creater | bigint | not null | - | ref users.uid | 创建者id |

---

*resources* - 资源

|字段|类型|null mode|默认值|约束|注释|
|:-:|:-:|:-:|:-:|:-:|:-:|
| rid | bigint | not null | - | - | 资源唯一标识 |
| rname | varchar(255) | not null | - | - | 资源名称 |
| creater | bigint | not null | - | ref users.uid | 创建者id |

---

*permissions* - 权限(操作-资源关联)

|字段|类型|null mode|默认值|约束|注释|
|:-:|:-:|:-:|:-:|:-:|:-:|
| pid | bigint | not null | - | - | 权限唯一标识 |
| action_id | bigint | not null | - | ref actions.aid | 操作id |
| action | varchar(64) | not null | - | ref actions.aname | 操作 |
| resource_id | bigint | not null | - | ref resources.rid | 资源id |
| resource | varchar(255) | not null | - | ref resources.rname | 资源 |
| creater | bigint | not null | - | ref users.uid | 创建者id |

---

*roles_to_permissions* - 角色-权限关联

|字段|类型|null mode|默认值|约束|注释|
|:-:|:-:|:-:|:-:|:-:|:-:|
| role_id | bigint | not null | - | ref roles.rid | 角色id |
| permission_id | bigint | not null | - | ref permission.pid | 权限id |
| creater | bigint | not null | - | ref users.uid | 创建者id |

**访问对象设计**

考虑如下现实:
- 数据库查询通常需要一个结构的字段作为匹配. 例如: User 在数据库中查询时依赖 User.Uid.
- 小对象或者基本类型复制的操作速度极快, 因此优先考虑栈分配.
- 尽可能保持一致的调用形式有利于统一处理.

基于上述现实, 设计如下访问对象:
- 结构对象统一后缀: _Dao, 通常嵌入 SqlMeta 结构.
- 对于允许为 null 的字段, 使用 sql.Nullxxx 类型适配.
- 所有字段需要显式使用字段标注来映射到db列.
- 跨表复杂对象构建不涉及列重构的情况下不单独映射dao结构.
- 对于序列返回结构通常引入 pagenation 结构.

对于repo层, 设计如下访问操作:
- 所有注入函数公共前缀: Set_
- 接口导出操作, 接口声明不需要形参声明.
- 接口定义需要形参声明, 规则为: 有且唯一结构总是为 dao; 其余按照语义进行声明; 必然返回 error; 总是具名返回.
- 简单结构值有用采用值类型, 复杂结构值有用采用指针类型; 简单结构值没用采用值类型, 复杂结构值没用采用值类型.
- 对于操作序列性更新库操作, 视情况开启事务.


一个简单的示例:
```go
// 对于 User 结构, 如果结构简单, User 或者 *User 均可
// 对于操作可能变化原有结构的使用指针传递, 否则使用值传递.
func InsertUser(dao *User) (err error) {}
func UpdateUser(dao *User) (err error) {}
// 查询需要指明查询匹配字段
// 这里 User 字段足够简单, 返回值即可
func QueryUserById(uid int) (dao User, err error) {}
func QueryUserByEmail(email string) (dao User, err error) {}
// 删除一般需要指明删除依据, 否则为直接关联id删除, 第一参数永远为 strict bool 用于控制软删除
func DeleteUser(strict bool, uid int) (err error) {}
func DeleteUserByEmail(strict bool, email string) (err error) {}
// 复杂关联操作需要考虑事务
func InsertUsers(dao []*User) (uids []int, err error) {}
func UpdateUsers(dao []*User) (uids []int, err error) {}
func QueryUsersByIds(uids []int) (dao Pagenation[*User], err error) {}
func QueryUsersByEmail(emails []string) (dao Pagenation[*User], err error) {}
func DeleteUsersByIds(strict bool, uids []int) (uids []int, err error) {}
func DeleteUsersByEmails(strict bool, emails []string) (uids []int, err error) {}
```

### <a id="design-peer">端点设计</a>

#### 用户

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

#### 角色

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

### <a id="design-config">配置设计</a>

## <a id="func-index">函数索引</a>

## <a id="client-sdk">客户端sdk</a>