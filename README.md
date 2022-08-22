# QQbot

#### 环境依赖
>go-cqhttp v1.0.0-rc3

#### 部署步骤
>1.安装go-cqhttp并做好相关调试  
>2.获取rasa机器人存入QQbot-rasa/models文件夹下并设置接口5005(默认)  
>3.设置config/config.yml文件的参数  
>4.设置global/global.go中**CfgFileURL** **MYQQID**字段的值  
>5.rasa train + rasa run 运行rasa机器人  
>6.运行./cmd/main.go

#### 目录结构描述
>├── README.md                   // help  
├── api                         // 路由层  
│  
├── cmd  
│　　└── main.go                 // main函数  
│  
├── config   
│　　├── config.go               // 配置控制   
│　　├── config.yml              // 配置文件  
│　　└── config_test.go  
│  
├── global  
│　　├── CQcode.go               // QQ表情信息&包装   
│　　├── message.go              // 描述通道传递的信息  
│　　├── types.go                // 全局结构体  
│　　└── vars.go                 // 全局常量和变量  
│  
├── pool                        // 协程连接池  
│　　├── pool.go                 // 协程预加载加载  
│　　└── routing_runtime.go      // 协程运行逻辑  
│  
├── QQbot-rasa                  // 储存rasa机器人模型  
│  
├── service                     // 服务层  
│　　└── respond.go              // 路由中间件  
│  
├── tools    
│　　├── rasa_tool               // 处理链接rasa机器人  
│　　│　　├── talk_test.go  
│　　│　　└── tool.go  
│　　│  
│　　├── routing_tool            // 协程维护层   
│　　│　　└── respond.go  
│　　│  
│　　├── server_tool             // 服务决断  
│　　│　　├── answer_post.go      // 回复层  
│　　│　　├── judge.go            // 判断层  
│　　│　　├── logic.go            // 回复逻辑层  
│　　│　　└── msg.go              // 信息处理层

#### 相关功能
