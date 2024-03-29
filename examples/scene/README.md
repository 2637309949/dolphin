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
    middles.xml    // 自定义中间件
    xlsx.xml       // excel解析

## Preparation

1. 启动事务管理器dtmsrv
```sh
$ git clone https://github.com/yedf/dtm
$ cp conf.sample.yml conf.yml
$ go run app/main.go dtmsvr
```

2. 启动kafka
```shell
$ docker run -d --name zookeeper -p 2181:2181 -t wurstmeister/zookeeper
$ docker run -d --name kafka \
-p 9092:9092 \
-e KAFKA_BROKER_ID=0 \
-e KAFKA_ZOOKEEPER_CONNECT=10.0.0.101:2181 \
-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://10.0.0.101:9092 \
-e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 wurstmeister/kafka
```

3. 启动redis

```yml
version: '3'
 
services:
 redis1:
  image: publicisworldwide/redis-cluster
  restart: always
  volumes:
   - /home/double/App/docker/redis/8001/data:/data
  environment:
   - REDIS_PORT=8001
  ports:
    - '8001:8001'     #服务端口
    - '18001:18001'   #集群端口
 
 redis2:
  image: publicisworldwide/redis-cluster
  restart: always
  volumes:
   - /home/double/App/docker/redis/8002/data:/data
  environment:
   - REDIS_PORT=8002
  ports:
    - '8002:8002'
    - '18002:18002'
 
 redis3:
  image: publicisworldwide/redis-cluster
  restart: always
  volumes:
   - /home/double/App/docker/redis/8003/data:/data
  environment:
   - REDIS_PORT=8003
  ports:
    - '8003:8003'
    - '18003:18003'
 
 redis4:
  image: publicisworldwide/redis-cluster
  restart: always
  volumes:
   - /home/double/App/docker/redis/8004/data:/data
  environment:
   - REDIS_PORT=8004
  ports:
    - '8004:8004'
    - '18004:18004'
 
 redis5:
  image: publicisworldwide/redis-cluster
  restart: always
  volumes:
   - /home/double/App/docker/redis/8005/data:/data
  environment:
   - REDIS_PORT=8005
  ports:
    - '8005:8005'
    - '18005:18005'
 
 redis6:
  image: publicisworldwide/redis-cluster
  restart: always
  volumes:
   - /home/double/App/docker/redis/8006/data:/data
  environment:
   - REDIS_PORT=8006
  ports:
    - '8006:8006'
    - '18006:18006'
```

启动redis
```shell
$ docker-compose  -f redis.yml  up -d
```

创建集群
```shell
$ docker run --rm -it inem0o/redis-trib \
create --replicas 1 \
172.20.10.3:8001 172.20.10.3:8002 172.20.10.3:8003 172.20.10.3:8004 172.20.10.3:8005 172.20.10.3:8006
```

4. 启动nsq
nsq 主要有三个组件: nsqlookupd, nsqd, nsqadmin。这三个组件都包含在 nsqio/nsq 镜像中

```shell
$ docker run --name lookupd -p 4160:4160 -p 4161:4161 -d nsqio/nsq /nsqlookupd
$ docker inspect -f '{{ .NetworkSettings.IPAddress }}' lookupd
$ docker run --name nsqd -p 4150:4150 -p 4151:4151 -p 4152:4152 -p 4153:4153  -d nsqio/nsq /nsqd \
--broadcast-address=172.20.10.3 \
--lookupd-tcp-address=172.20.10.3:4160 \
-tcp-address=0.0.0.0:4152 \
-http-address=0.0.0.0:4153
$ docker run -d --name nsqadmin -p 4171:4171 nsqio/nsq /nsqadmin \
--lookupd-http-address=172.20.10.3:4161
```
