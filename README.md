
casbin体验

### 目录结构
```
├── main.go             //程序入口
├── go.mod              //go的包管理
├── README.md           
├── config                //配置文件等
│   ├── rbac_model.conf
├── global                //全局变量
│   ├── global.go
├── middleware          //中间件，casbin权限等
├── module             //功能模块
    ├── model          //请求，返回等结构体
    ├── router         //接口路由
    │   └── ...
    ├── api             //具体api的方法
    ├── service         //操作数据库方法，dao层
```

### 关键理解点
- ParamsMatchFunc函数如何注入rbac_model.conf中(model/casbin.go)
- Casbin Model 的语法和运行逻辑