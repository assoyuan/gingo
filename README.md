- 介绍

```
基于Gin封装的一个轻量级MCV项目结构，便于使用者快速进入业务开发。

集成的组件包括：

配置 [https://github.com/spf13/viper]
日志 [https://github.com/sirupsen/logrus]
日志分割 [https://github.com/natefinch/lumberjack]
ORM(待完善)

持续更新中...
```

- 目录介绍

```
.
├── config                  // 配置文件目录
│   └── config.go           // 读取配置文件代码文件
├── controllers             // 控制器目录
│   ├── auth.go             // 处理登录、注册等控制器的代码文件
│   └── user.go             // 处理用户 CRUD 操作的控制器的代码文件
├── middlewares             // 中间件目录
│   ├── auth.go             // 处理用户鉴权的中间件代码文件
│   └── logger.go           // 日志记录中间件代码文件
├── models                  // 数据库模型目录
│   ├── db.go               // 初始化数据库连接代码文件
│   ├── user.go             // User 模型代码文件
│   └── ...                 // 其他模型代码文件
├── routes                  // 路由目录
│   ├── auth.go             // 登录、注册等路由代码文件
│   └── user.go             // 用户管理相关路由代码文件
├── static                  // 静态资源目录
│   ├── css                 // 样式目录
│   ├── js                  // JS 文件目录
│   └── ...                 // 其他资源目录
├── templates               // HTML 模板目录
│   ├── base.html           // 基本模板文件
│   ├── auth                // 登录、注册相关模板目录
│   ├── user                // 用户相关模板目录
│   └── ...                 // 其他模板目录
├── tests                   // 测试文件目录
│   ├── controllers_test.go // 控制器测试代码文件
│   ├── middlewares_test.go // 中间件测试代码文件
│   └── models_test.go      // 模型测试代码文件
├── utils                   // 工具函数目录
│   ├── jwt.go              // JWT 相关代码文件
│   ├── response.go         // 响应处理相关代码文件
│   └── ...                 // 其他工具代码文件
├── main.go                 // 项目入口文件
├── config.yaml             // 配置文件
├── go.mod                  // Go 依赖管理文件
├── go.sum                  // Go 依赖管理锁定文件
└── README.md               // 项目说明文档

这是一个典型的 MVC（模型-视图-控制器）项目结构，其中：

main.go     项目的入口文件。
config      目录是存放项目的配置文件的目录，包含了读取配置文件的代码文件 config.go。
controllers 目录是存放控制器的目录，包含了处理不同业务逻辑的控制器代码文件。
middlewares 目录是存放中间件的目录，包含了处理请求过程中的各种中间件代码文件。
models      目录是存放数据库模型的目录，包含了建立数据库连接和定义数据库模型的代码文件。
routes      目录是存放路由配置文件的目录，包含了各个业务路由的代码文件。
static      目录是存放静态资源的目录，例如 CSS、JS、图片等。
templates   目录是存放 HTML 模板文件的目录，用于渲染视图。
tests       目录是存放测试文件的目录，包含了控制器、中间件和模型的测试代码。
utils       目录是存放工具函数的目录，包含了各种辅助函数，例如 JWT 处理、响应处理等。
```

- 启动

```
#编译运行
go run main.go

#打包运行
go build -o gingo main.go
./gingo
```