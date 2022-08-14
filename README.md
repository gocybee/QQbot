# QQbot

#### 环境依赖
>go-cqhttp v1.0.0-rc3  
>mysql  

#### 部署步骤
>安装go-cqhttp  
>设置config/config.yml文件的参数  
>设置global/global.go中**SendMsgURL** **CfgFileURL** **MYQQID**字段的值  
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
│   └── respond.go              // 路由中间件  
│  
├── tools                       // 工具层  
│   ├── daoTools.go             // 数据库信息操作   
│   ├── serviceTools.go         // 服务决断  
│   └── tools.go                // 常用工具

#### 相关功能
>1.私聊时可以进行学习，导出文件，正常问答的问题  
> &emsp;需要学习：回复”问题1 问题2 问题3+答案“即可收到回复  
> &emsp;需要导出文件：回复”导出问答文件“即生成名字为oldMsg.yml文件  
>2.注意：所有学习的语句将会被模糊查询，并在所有情景下回答。