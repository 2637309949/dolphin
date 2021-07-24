## Application scenario example

    ami.xml      // 基于redis的队列处理  
    article.xml  // CURD生成代码  
    cache.xml    // 缓存中间件配置  
    encrypt.xml  // 加密中间件测试  
    kafka.xml    // kafka用例  
    redis.xml    // redis分布式锁  
    rpc.xml      // grpc跨进程调用  
    sqlmap.xml   // sql文件加载  
    user.xml     // 用户信息加载  
    view.xml     // view视图函数  
    dtm.xml      // 分布式事务

## Preparation

1. 启动事务管理器dtmsrv
```sh
git clone https://github.com/yedf/dtm
cp conf.sample.yml conf.yml # 修改conf.yml， 执行dtmsvr/dtmsvr.sql
go run app/main.go dtmsvr
INFO[0000] start dtmsvr                                 
2021-07-23 17:44:24.371 main.go:27 dtmsvr listen at: 8080
2021-07-23 17:44:24.474 types.go:117 connecting root:****@tcp(localhost:3306)/dtm?charset=utf8mb4&parseTime=true&loc=Local
```

2. 启动kafka

3. 启动redis

4. 启动nsq


