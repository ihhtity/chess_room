# 棋牌室多端系统 - 功能规格文档

## 一、项目概述

本系统是一个适配微信小程序、微信H5、普通H5的多端棋牌室预订管理系统，分为用户端和管理端。

### 技术栈

| 层级 | 技术 | 说明 |
|------|------|------|
| 后端 | Go + Gin | RESTful API 服务 |
| 数据库 | MySQL | 主业务数据存储 |
| 缓存 | Redis | 会话管理、热点数据缓存 |
| 用户端 | React + TypeScript + Vite | Web端用户界面框架 |
| 用户端 | NutUI React | 移动端 UI 组件库 |
| 用户端 | axios | HTTP 请求库 |
| 用户端 | react-router-dom | 路由管理 |
| 用户端 | Sass | CSS 预处理器 |
| 管理端 | React + TypeScript + Vite | PC端管理后台框架 |
| 管理端 | Ant Design | 企业级 UI 组件库 |
| 管理端 | axios | HTTP 请求库 |
| 管理端 | react-router-dom | 路由管理 |

### 开发脚本

**用户端 (frontend)**

| 命令 | 说明 |
|------|------|
| `npm run dev` | 启动开发服务器 |
| `npm run build` | 构建生产版本 |
| `npm run preview` | 预览生产构建 |
| `npm run lint` | 执行 ESLint 代码检查 |

**管理端 (admin)**

| 命令 | 说明 |
|------|------|
| `npm run dev` | 启动开发服务器（端口 3001） |
| `npm run build` | 构建生产版本 |
| `npm run preview` | 预览生产构建 |

### 多端适配

- **Web端**: 浏览器访问，支持手机号登录、在线支付
- **管理端**: PC浏览器访问，支持管理员登录、后台管理

---

## 二、用户端功能

### 2.1 用户认证

- 微信登录（H5）
- 手机号验证码登录
- 注册/绑定手机号
- 用户信息编辑

### 2.2 包间浏览

- 包间列表（按类型、状态筛选）
- 包间详情（设备、图片、价格）
- 实时状态展示（空闲/使用中/已预约）

### 2.3 预约系统

- 选择日期和时段
- 查看时段价格
- 预约确认
- 预约取消（提前2小时免费取消）

### 2.4 订单管理

- 订单列表（待支付、使用中、已完成、已取消）
- 订单详情
- 在线支付（微信支付）
- 支付记录

### 2.5 会员系统

- 会员等级（普通/白银/黄金/钻石）
- 储值功能
- 会员折扣
- 消费积分

### 2.6 消息通知

- 预约成功通知
- 支付成功通知
- 开场提醒
- 结束提醒

---

## 三、管理端功能

### 3.1 管理员认证

- 管理员登录
- 权限管理（超级管理员/普通管理员）
- 操作日志

### 3.2 包间管理

- 包间列表 CRUD
- 包间类型管理
- 包间状态管理（维护/启用）
- 设备管理

### 3.3 计费规则

- 时段设置（早场/午场/晚场/通宵）
- 价格规则（按时段、按包间类型、会员价）
- 折扣规则
- 包场规则

### 3.4 订单管理

- 订单列表
- 订单状态变更
- 手动开单
- 订单统计

### 3.5 会员管理

- 会员列表
- 会员等级调整
- 储值记录
- 积分管理

### 3.6 数据统计

- 日/周/月营业报表
- 包间使用率统计
- 收入分析
- 用户增长统计

### 3.7 系统设置

- 营业时间设置
- 支付配置
- 微信配置
- 通知模板配置

---

## 四、核心业务流程

### 4.1 预约流程

```
浏览包间 → 选择日期时段 → 确认订单 → 在线支付 → 预约成功 → 开场提醒 → 入场使用 → 结束结算
```

### 4.2 即时开单流程

```
选择包间 → 选择起始时间 → 创建订单 → 支付 → 开始计时 → 结束计时 → 结算
```

### 4.3 会员储值流程

```
进入会员中心 → 选择储值金额 → 在线支付 → 储值成功 → 余额增加
```

---

## 五、数据库设计

### 5.1 核心数据表

| 表名 | 说明 |
|------|------|
| users | 用户信息表 |
| admins | 管理员表 |
| room_types | 包间类型表 |
| rooms | 包间信息表 |
| time_slots | 时段价格表 |
| orders | 订单表 |
| payments | 支付流水表 |
| memberships | 会员表 |
| recharge_records | 储值记录表 |
| notifications | 消息通知表 |
| holidays | 节假日表 |
| operation_logs | 操作日志表 |

### 5.2 字段设计概览

#### users 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键 |
| openid | VARCHAR(100) | 微信openid |
| phone | VARCHAR(20) | 手机号 |
| password | VARCHAR(255) | 密码（bcrypt加密） |
| nickname | VARCHAR(50) | 昵称 |
| realname | VARCHAR(50) | 真实姓名 |
| avatar | VARCHAR(255) | 头像 |
| gender | TINYINT | 性别（0未知/1男/2女） |
| status | TINYINT | 状态（0禁用/1启用） |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

#### admins 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键 |
| username | VARCHAR(50) | 用户名 |
| password | VARCHAR(255) | 密码（bcrypt加密） |
| realname | VARCHAR(50) | 真实姓名 |
| role | TINYINT | 角色（0超级管理员/1普通管理员） |
| status | TINYINT | 状态（0禁用/1启用） |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

#### room_types 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键 |
| name | VARCHAR(50) | 类型名称 |
| description | TEXT | 描述 |
| base_price | DECIMAL(10,2) | 基础价格 |
| max_people | INT | 最大容纳人数 |
| sort_order | INT | 排序 |
| status | TINYINT | 状态（0禁用/1启用） |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

#### rooms 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键 |
| name | VARCHAR(50) | 包间名称 |
| type_id | BIGINT | 包间类型ID |
| floor | VARCHAR(20) | 楼层 |
| capacity | INT | 容量 |
| equipment | JSON | 设备清单 |
| images | JSON | 图片列表 |
| description | TEXT | 描述 |
| status | TINYINT | 状态（0维护/1启用） |
| sort_order | INT | 排序 |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

#### time_slots 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键 |
| type_id | BIGINT | 包间类型ID |
| name | VARCHAR(50) | 时段名称 |
| start_time | TIME | 开始时间 |
| end_time | TIME | 结束时间 |
| price | DECIMAL(10,2) | 基础价格 |
| weekday_price | DECIMAL(10,2) | 工作日价格 |
| weekend_price | DECIMAL(10,2) | 周末价格 |
| holiday_price | DECIMAL(10,2) | 节假日价格 |
| sort_order | INT | 排序 |
| status | TINYINT | 状态（0禁用/1启用） |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

#### orders 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键 |
| order_no | VARCHAR(32) | 订单号 |
| user_id | BIGINT | 用户ID |
| room_id | BIGINT | 包间ID |
| start_time | DATETIME | 开始时间 |
| end_time | DATETIME | 结束时间 |
| actual_end_time | DATETIME | 实际结束时间 |
| duration | INT | 时长（分钟） |
| status | TINYINT | 状态（0待支付/1使用中/2已完成/3已取消） |
| total_amount | DECIMAL(10,2) | 总金额 |
| paid_amount | DECIMAL(10,2) | 已支付金额 |
| paid_at | DATETIME | 支付时间 |
| completed_at | DATETIME | 完成时间 |
| cancel_time | DATETIME | 取消时间 |
| remark | TEXT | 备注 |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

#### payments 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键 |
| order_id | BIGINT | 订单ID |
| user_id | BIGINT | 用户ID |
| amount | DECIMAL(10,2) | 支付金额 |
| payment_type | TINYINT | 支付类型（1微信支付/2支付宝/3余额支付） |
| status | TINYINT | 状态（0待支付/1成功/2失败） |
| transaction_no | VARCHAR(64) | 交易号 |
| paid_at | DATETIME | 支付时间 |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

#### memberships 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGINT | 主键 |
| user_id | BIGINT | 用户ID |
| level | TINYINT | 会员等级（0普通/1白银/2黄金/3钻石） |
| points | INT | 积分 |
| balance | DECIMAL(10,2) | 余额 |
| total_consumed | DECIMAL(10,2) | 累计消费 |
| total_recharged | DECIMAL(10,2) | 累计储值 |
| membership_status | TINYINT | 状态（0禁用/1启用） |
| joined_at | DATETIME | 入会时间 |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

---

## 六、API 设计

### 6.1 用户端 API

| 模块 | 路径 | 方法 | 说明 |
|------|------|------|------|
| 认证 | /api/user/login | POST | 用户登录 |
| 认证 | /api/user/register | POST | 用户注册 |
| 认证 | /api/user/info | GET | 获取用户信息 |
| 认证 | /api/user | PUT | 更新用户信息 |
| 认证 | /api/user/password | PUT | 修改密码 |
| 包间 | /api/rooms | GET | 获取包间列表 |
| 包间 | /api/rooms/:id | GET | 获取包间详情 |
| 包间 | /api/rooms/:id/availability | GET | 获取包间可用时段 |
| 预约 | /api/orders | POST | 创建订单 |
| 预约 | /api/orders | GET | 获取订单列表 |
| 预约 | /api/orders/:id | GET | 获取订单详情 |
| 预约 | /api/orders/:id | PUT | 更新订单 |
| 预约 | /api/orders/:id/cancel | POST | 取消订单 |
| 支付 | /api/payments | POST | 创建支付 |
| 支付 | /api/payments/:id | GET | 获取支付详情 |
| 支付 | /api/payments | GET | 获取支付列表 |
| 支付 | /api/payments/notify | POST | 支付回调 |
| 支付 | /api/payments/order/:orderNo | GET | 根据订单号查询支付 |
| 支付 | /api/payments/wechat | POST | 微信支付 |
| 会员 | /api/membership | GET | 获取会员信息 |
| 会员 | /api/recharge | POST | 储值 |
| 会员 | /api/recharge/records | GET | 获取储值记录 |
| 消息 | /api/notifications | GET | 获取通知列表 |
| 微信 | /api/wechat/login | POST | 微信登录 |
| 微信 | /api/wechat/pay | POST | 微信支付 |

### 6.2 管理端 API

| 模块 | 路径 | 方法 | 说明 |
|------|------|------|------|
| 认证 | /api/admin/login | POST | 管理员登录 |
| 认证 | /api/admin/info | GET | 获取管理员信息 |
| 包间 | /api/admin/rooms | POST | 创建包间 |
| 包间 | /api/admin/rooms | GET | 获取包间列表 |
| 包间 | /api/admin/rooms/:id | GET | 获取包间详情 |
| 包间 | /api/admin/rooms/:id | PUT | 更新包间 |
| 包间 | /api/admin/rooms/:id | DELETE | 删除包间 |
| 类型 | /api/admin/room-types | POST | 创建包间类型 |
| 类型 | /api/admin/room-types | GET | 获取包间类型列表 |
| 类型 | /api/admin/room-types/:id | PUT | 更新包间类型 |
| 类型 | /api/admin/room-types/:id | DELETE | 删除包间类型 |
| 时段 | /api/admin/time-slots | POST | 创建时段 |
| 时段 | /api/admin/time-slots | GET | 获取时段列表 |
| 时段 | /api/admin/time-slots/:id | PUT | 更新时段 |
| 时段 | /api/admin/time-slots/:id | DELETE | 删除时段 |
| 订单 | /api/admin/orders | GET | 订单列表 |
| 订单 | /api/admin/orders/:id | GET | 订单详情 |
| 订单 | /api/admin/orders/:id | PUT | 更新订单状态 |
| 订单 | /api/admin/orders/:id/checkin | POST | 订单入场 |
| 订单 | /api/admin/orders/:id/checkout | POST | 订单结算 |
| 会员 | /api/admin/memberships | GET | 会员列表 |
| 会员 | /api/admin/memberships/:id | GET | 会员详情 |
| 会员 | /api/admin/memberships/:id | PUT | 更新会员信息 |
| 用户 | /api/admin/users | GET | 用户列表 |
| 用户 | /api/admin/users/:id | GET | 用户详情 |
| 用户 | /api/admin/users/:id | PUT | 更新用户状态 |
| 统计 | /api/admin/statistics | GET | 数据统计 |
| 日志 | /api/admin/logs | GET | 操作日志 |
| 节假日 | /api/admin/holidays | POST | 创建节假日 |
| 节假日 | /api/admin/holidays | GET | 获取节假日列表 |
| 节假日 | /api/admin/holidays/:id | DELETE | 删除节假日 |
| 管理员 | /api/admin/admins | POST | 创建管理员 |
| 管理员 | /api/admin/admins | GET | 获取管理员列表 |
| 管理员 | /api/admin/admins/:id | PUT | 更新管理员 |
| 管理员 | /api/admin/admins/:id | DELETE | 删除管理员 |

---

## 七、部署方案

### 7.1 目录结构

```
chess-room/
├── backend/                  # Go 后端
│   ├── config/               # 配置文件
│   │   └── config.go         # 配置读取
│   ├── controller/           # 控制器
│   │   ├── admin.go          # 管理员控制器
│   │   ├── room.go           # 包间控制器
│   │   ├── order.go          # 订单控制器
│   │   ├── payment.go        # 支付控制器
│   │   ├── membership.go     # 会员控制器
│   │   ├── user.go           # 用户控制器
│   │   └── wechat.go         # 微信相关控制器
│   ├── dao/                  # 数据访问层
│   │   └── mysql/            # MySQL数据访问
│   │       ├── admin.go
│   │       ├── room.go
│   │       ├── order.go
│   │       ├── payment.go
│   │       ├── membership.go
│   │       ├── user.go
│   │       └── mysql.go      # 数据库连接
│   ├── logic/                # 业务逻辑层
│   │   ├── admin.go
│   │   ├── room.go
│   │   ├── order.go
│   │   ├── payment.go
│   │   ├── membership.go
│   │   └── user.go
│   ├── model/                # 数据模型
│   │   ├── model.go          # 模型定义
│   │   └── seed.go           # 初始化数据
│   ├── middleware/           # 中间件
│   │   ├── auth.go           # JWT认证
│   │   ├── cors.go           # 跨域处理
│   │   ├── logger.go         # 日志记录
│   │   └── recovery.go       # 错误恢复
│   ├── router/               # 路由
│   │   └── router.go         # 路由配置
│   ├── pkg/                  # 工具包
│   │   ├── errno/            # 错误码
│   │   ├── log/              # 日志
│   │   ├── response/         # 统一响应
│   │   ├── utils/            # 工具函数
│   │   └── wechat/           # 微信SDK封装
│   ├── main.go               # 入口
│   └── go.mod                # 依赖
├── frontend/                 # React 用户端
│   ├── src/
│   │   ├── api/              # API接口定义
│   │   ├── components/       # 组件
│   │   ├── pages/            # 页面
│   │   │   ├── booking/      # 预约相关
│   │   │   ├── order/        # 订单相关
│   │   │   ├── member/       # 会员相关
│   │   │   ├── room/         # 包间相关
│   │   │   └── user/         # 用户相关
│   │   ├── router/           # 路由配置
│   │   ├── utils/            # 工具函数
│   │   ├── app.tsx           # 应用根组件
│   │   ├── main.tsx          # 入口
│   │   └── app.scss          # 全局样式
│   ├── vite.config.ts        # Vite配置
│   ├── tsconfig.json         # TypeScript配置
│   └── package.json          # 依赖配置
├── admin/                    # React 管理端
│   ├── src/
│   │   ├── api/              # API接口定义
│   │   ├── components/       # 组件
│   │   ├── pages/            # 页面
│   │   ├── router/           # 路由配置
│   │   ├── utils/            # 工具函数
│   │   ├── app.tsx           # 应用根组件
│   │   ├── main.tsx          # 入口
│   │   └── app.scss          # 全局样式
│   ├── vite.config.ts        # Vite配置
│   ├── tsconfig.json         # TypeScript配置
│   └── package.json          # 依赖配置
└── PROJECT_SPEC.md           # 项目规格文档
```

### 7.2 环境变量

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=chess_room

REDIS_HOST=localhost
REDIS_PORT=6379

JWT_SECRET=your_jwt_secret_here
WECHAT_APP_ID=your_wechat_app_id
WECHAT_APP_SECRET=your_wechat_app_secret
WECHAT_MCH_ID=your_wechat_mch_id
WECHAT_API_KEY=your_wechat_api_key

SERVER_PORT=8080
```

### 7.3 数据库自动初始化

系统启动时会自动执行以下操作：

1. **数据库检查与创建**: 如果数据库不存在，自动创建
2. **数据表迁移**: 使用 GORM AutoMigrate 自动创建/更新表结构
3. **默认数据初始化**: 创建默认管理员、包间类型、时段等基础数据
4. **测试数据初始化**: 创建测试用户、订单、会员等模拟数据

---

## 八、测试数据说明

系统启动时会自动创建以下模拟测试数据：

### 8.1 默认管理员

| 用户名 | 密码 | 角色 |
|--------|------|------|
| admin | 123456 | 超级管理员 |

### 8.2 包间类型

| ID | 名称 | 价格 | 最大人数 |
|----|------|------|----------|
| 1 | 普通棋牌室 | 50元 | 4人 |
| 2 | 豪华棋牌室 | 80元 | 6人 |
| 3 | VIP包间 | 120元 | 8人 |

### 8.3 包间

| ID | 名称 | 类型 | 楼层 | 状态 |
|----|------|------|------|------|
| 1 | 101室 | 普通 | 1层 | 启用 |
| 2 | 102室 | 普通 | 1层 | 启用 |
| 3 | 201室 | 豪华 | 2层 | 启用 |
| 4 | 202室 | 豪华 | 2层 | 启用 |
| 5 | 301室 | VIP | 3层 | 启用 |

### 8.4 测试用户

| ID | 昵称 | 手机号 | 密码 |
|----|------|--------|------|
| 1 | 张三 | 13800138001 | 123456 |
| 2 | 李四 | 13800138002 | 123456 |
| 3 | 王五 | 13800138003 | 123456 |
| 4 | 赵六 | 13800138004 | 123456 |
| 5 | 钱七 | 13800138005 | 123456 |

### 8.5 测试订单

| ID | 用户 | 状态 | 金额 |
|----|------|------|------|
| 1 | 张三 | 待支付 | 150元 |
| 2 | 李四 | 使用中 | 250元 |
| 3 | 王五 | 已完成 | 400元 |
| 4 | 赵六 | 已取消 | 240元 |

---

## 九、开发计划

### 阶段一：基础架构 ✅

1. 搭建后端项目骨架
2. 配置 MySQL 和 Redis
3. 实现基础中间件（CORS、JWT、日志）

### 阶段二：核心功能 ✅

1. 包间管理模块
2. 预约订单模块
3. 支付模块
4. 会员模块

### 阶段三：前端开发 ✅

1. 用户端 React + Vite + NutUI 开发
2. 管理端 React + Vite + Ant Design 开发

### 阶段四：集成测试

1. 微信登录集成
2. 微信支付集成
3. 消息通知集成
4. 性能测试

---

## 十、注意事项

1. 所有敏感配置通过环境变量注入
2. 数据库操作使用参数化查询防止 SQL 注入
3. 接口需进行输入校验
4. 支付回调需进行签名验证
5. 订单状态需考虑并发问题，使用 Redis 分布式锁确保一致性
6. 密码使用 bcrypt 加密存储
7. API 响应使用统一格式
8. 跨域请求需配置 CORS 中间件，允许前端和管理端访问
