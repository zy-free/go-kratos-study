- 为什么包取名叫internal，因为go会对这个包名做禁止引用，保证代码的解耦

- B站v1的httpServer运行
    - [x] log
    - [x] maxbyte
    - [x] maxconn
        - 注意limiter初始化的位置 
    - [x] recover
    - [x] timeout（http层次未作拦截，即无传入timeoutHander,但是传入了Context供其他层判断）
    - [ ] trace(jaeger)
    - [ ] limit(根据cpu过载保护，https://github.com/alibaba/Sentinel/wiki/%E7%B3%BB%E7%BB%9F%E8%87%AA%E9%80%82%E5%BA%94%E9%99%90%E6%B5%81)
        - 压测
    - [ ] 分布式限流(不是简单的redis分布式限流，会有热key瓶颈)
    - [ ] criticality 接口重要性
    - [ ] promethues监控
    - [x] 新增break中间件(接口级别熔断)
    - [ ] restful路由支持(:id，path参数支持重复)
    - [x] default标签改写，之前json格式时不生效


    
- grpc功能实现
    - [x] etcd注册
    - [x] p2c负载均衡
    - [x] break
    - [x] log
    - [x] grpcerror
    - [x] metadata
    - [x] 级联timeout
    - [x] trace
    
- demo
    - [ ] admin+gorm(share db架构,运营平台微服务共享db，只是权限不同)
    - [x] service+mysql
    - [x] job+kafka
    - [ ] canal,发送kafka,异步删除缓存
    - [ ] Beanstalkd?
    - [x] kafka
    - [x] 数据库delete_time 为null的时间处理，自定义时间
    - [x] errgroup
    - [x] runsafe
    - [x] 获取天气的等api的demo，client以及分层设计
    - [x] mysql慢日志，熔断
    - [x] redis慢日志，无熔断(不支持集群)
    - [x] redis的demo，读失败后的写缓存策略（降级后一般读失败不触发回写缓存）。
    - [x] 空缓存保护策略,将空数据缓存，避免请求直接打到db
    - [x] redis分布式锁(架构上尽量规避，性能不高，case多，容易出bug)
    - [ ] 基于B站的mysql的redis封装go-zero的cache
    - [x] chan-singleFlight
    - [x] chan-fanout(生产消费模式：这里用作redis回写异步处理)
    - [x] chan-pipeline
    - [x] metadata
    - [x] grpcerror
    - [x] validate
    - [x] 贫血模型
    - [x] csv导出
    - [x] breaker(google sre 熔断器)
    - [x] 多租户 （流量染色）
    - [x] hash-id封装(内部int类型给前端时转化为hash)
    - [x] attrs(mysql字段，bit标识)
    - [x] redis(bitmap大offset处理)
    - [ ] 滑动窗口

    
- [x]  pkg/errors处理指南
    - 底层包不需要Wrap，直接返回，或者定义sentinel error
    - dao层Wrap，或者其他层第一次定义错误码，也要Wrap
    - 已经wrap的err，避免堆栈信息重复，可使用WtihMessage带信息，不处理可直接return
    - 在最外层 %+v 打印堆栈信息
    - 个人感觉在grpc的client端没必要加日志