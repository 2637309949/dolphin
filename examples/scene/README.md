## Application scenario example

    ami.xml        // 基于redis的队列处理  
    article.xml    // CURD生成代码  
    cache.xml      // 缓存中间件配置  
    dtm.xml        // 分布式事务
    encrypt.xml    // 加密中间件测试  
    jwt.xml        // JWT认证   
    kafka.xml      // kafka用例  
    nsq.xml        // NSQ消息发布订阅
    redislock.xml  // redis分布式锁 
    redismq.xml    // redis消息队列 
    rpc.xml        // grpc跨进程调用  
    sqlmap.xml     // sql文件加载  
    user.xml       // 用户信息加载  
    view.xml       // view视图函数  
    vote.xml       // vote投票

## Preparation

1. 启动事务管理器dtmsrv
```sh
$ git clone https://github.com/yedf/dtm
$ cp conf.sample.yml conf.yml
$ go run app/main.go dtmsvr
```

2. 启动kafka

3. 启动redis

4. 启动nsq


