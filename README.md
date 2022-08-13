# QQbot

#### 环境依赖
>go-cqhttp v1.0.0-rc3  
>mysql  

#### 部署步骤
>安装go-cqhttp  
>运行./cmd/main.go

#### 目录结构描述
├── README.md                   // help  
├── api                         // 路由层  
├── cmd                         // main函数  
├── config                      // 配置  
│   ├── config.go               // 配置控制   
│   ├── config.yml              // 配置文件  
│   └── config_test.go  
│
├── dao                         // 数据库操作层  
│   ├── dao.go                  //数据库操作  
│   └── select_test.go  
│
├── global                      // 全局变量层  
│   ├── CQcode.go               // 特殊信息包装   
│   └── global.go               // 全局常量和变量  
│
├── service                     // 服务层  
│   ├── respond.go              // 路由中间件   
│
├── tools                       // 工具层  
│   ├── daoTools.go             // 数据库信息操作   
│   ├── serviceTools.go         // 服务决断  
│   └── tools.go                // 常用工具