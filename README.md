- 为什么包取名叫internal，因为go会对这个包名做禁止引用，保证代码的解耦

- B站v1的httpServer运行
    - [x] log
    - [x] maxbyte
    - [x] maxconn
        - 注意limiter初始化的位置 
    - [x] recover
    - [x] timeout（http层次未作拦截，即无传入timeoutHander,但是传入了Context供其他层判断）
    - [ ] trace(zip)
    - [ ] limit(根据cpu自适应限流，https://github.com/alibaba/Sentinel/wiki/%E7%B3%BB%E7%BB%9F%E8%87%AA%E9%80%82%E5%BA%94%E9%99%90%E6%B5%81)
        - 压测
    - [ ] break(自适应熔断，google sre算法)
    - [ ] promethues监控

    
- grpc功能实现
    - [x] etcd注册
    - [x] p2c负载均衡
    - [ ] break
    - [x] log
    - [x] grpcerror
    - [x] metadata
    - [ ] 级联timeout
    - [x] trace
    
    
- demo
    - [ ] admin+gorm(share db架构,运营平台微服务共享db)
    - [ ] service+sql
    - [ ] job+Beanstalkd
    - [ ] 数据库delete_time 为null的时间处理，自定义时间
    - [ ] 多租户 （流量染色）
    - [ ] errgroup
    - [ ] mysql慢日志，熔断
    - [ ] redis慢日志，熔断
    - [ ] 获取天气的等api的demo，client以及分层设计
    - [x] metadata
    - [x] grpcerror
    - [ ] validate
    - [ ] timeout
    - [x] 贫血模型
    - [ ] 流量染色
    - [ ] csv导出
    
    
- [x]  pkg/errors处理指南
    - 底层包不需要Wrap，直接返回，或者定义sentinel error
    - dao层Wrap，或者其他层第一次定义错误码，也要Wrap
    - 已经wrap的err，避免堆栈信息重复，可使用WtihMessage带信息，不处理可直接return
    - 在最外层 %+v 打印堆栈信息
    - 个人感觉在grpc的client端没必要加日志