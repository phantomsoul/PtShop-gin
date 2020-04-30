# pt-Gin PtShop Server
> 基于[Gin](https://github.com/gin-gonic/gin)的商城项目服务端框架，为了大家熟练掌握gin框架的相关技能。欢迎coder积极提工单，打造优美·简洁·高效的代码，使项目越来越强大。o(*≧▽≦)ツ加油

使用gin构建了企业级商城脚手架，代码简洁易读，可快速进行高效API开发。
 主要功能有：
 1. [INI](https://github.com/go-ini/ini) 地表最强大、最方便和最流行的 Go 语言 INI 文件操作库。
 2. [Validator](https://github.com/go-playground/validator) 接入validator.v10，支持多语言错误信息提示及自定义错误提示。

## Installation

先编辑```server.ini、config/env.go```中相关配置，然后运行：
```
make
```

浏览器访问 http://localhost:3000/api/index

OS X & Linux:

```
cd pt-gin
go mod tidy
go mod vendor
```

## directory structure

```
.
├── Makefile
├── config                      全局配置
│   ├── conf.go                 主配置             
│   ├── jwt.go
│   └── env.go                  环境配置
│
├── connections                 存储连接
│   ├── database
│   │   └── mysql
│   └── redis
│       └── redis.go
│
├── controllers                 控制器
│   ├── controller.go           控制器初始化
│   └── SampleController.go     示例代码参考
│ 
├── middleware                  中间件
│   ├── auth                    认证中间件
│   │   ├── drivers             认证引擎
│   │   └── auth.go   
│   │            
│   ├── verify                  登录安全验证中间件
│   │   └── verify.go  
│   │ 
│   └── middleware.go    
│         
├── public                      静态资源
│   ├── assets
│   │   ├── css
│   │   ├── images
│   │   └── js
│   ├── dist
│   └── templates
│       └── index.tpl
├── handle.go                   全局错误处理
│
├── main.go                     入口 
│
├── models                      数据模型
│
├── dao                         数据操作
│   └── dao.go                  数据操作初始化
│
├── modules                     项目模块
│   │── ecode                   返回码封装
│   │   ├── ecode.go
│   │   └── err_code.go
│   │── pack                    返回JSON格式封装
│   │   └── pack.go    
│   │── types                   各组件配置文件
│   │   └── config.go
│   │── schedule
│   │   └── schedule.go         定时任务模块
│   │── log                     日志
│   │   └── log.go 
│   │── util                    公共方法封装
│   │   └── util.go 
│   │── xhttp                   HTTP外部访问方法封装
│   │   └── xhttp.go 
│   └── sig                     服务启停信号管理
│       └── sig.go           
├── routers                     路由
│   └── api_routers.go          
├── routers.go                  路由初始化设置
│
├── service                     服务(业务逻辑处理)
│   └── service.go              服务初始化           
│
├── build                       编译输出路径
│   └── log                     日志文件夹        
│
├── server.ini                  框架主配置文件
│
└── vendor                      govendor 第三方包
```

### HTTP 层(基于[Gin](https://github.com/gin-gonic/gin))
- 路由
- 中间件
- 控制器
- 请求
- 响应
- 视图
- JWT

### 安全
- 用户认证
- 用户授权
- 加密解密
- 哈希

### 综合
- dancer 命令行
- 缓存系统
- 错误与日志
- 任务调度

### 数据库
- mysql
- redis


## 项目依赖
- Web框架：github.com/gin-gonic/gin
- ORM：github.com/jinzhu/gorm
- JWT：github.com/dgrijalva/jwt-go
- Redis：github.com/garyburd/redigo/redis
- Mysql：github.com/go-sql-driver/mysql
- Log：go.uber.org/zap
- 任务调度：github.com/robfig/cron