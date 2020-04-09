# vehicle

用到的技术
gin、goredis、gorm、restfulapi、protobuf、emq、jwt、gomod

项目的总体构建流程如下:
![image](https://github.com/iyongfei/vehicle/blob/master/assets/outer_line.png)

仓促完成，后续在更新

```

.
├── README.md
├── assets.   存放sql文件
│   ├── gomod
│   ├── outer_line.png
│   └── vehicle.sql 数据库sql文件
├── conf.ini  项目的配置文件
├── conf.txt  测试脚本服务的配置文件
├── go.mod    mod配置文件
├── go.sum
├── src
│   ├── vehicle  项目服务目录
│   │   ├── api_server  API包
│   │   │   ├── auth.go 
│   │   │   ├── threat.go
│   │   │   ├── vehicle.go
│   │   │   └── white_list.go
│   │   ├── conf   ini库读取配置
│   │   │   ├── conf.go
│   │   │   └── conf_util.go
│   │   ├── cron   定时任务包
│   │   │   ├── cron.go
│   │   │   └── cron_func.go
│   │   ├── db     存储
│   │   │   ├── db.go
│   │   │   ├── mysql 
│   │   │   └── redis
│   │   ├── docs
│   │   │   ├── docs.go
│   │   │   ├── swagger.json
│   │   │   └── swagger.yaml
│   │   ├── emq
│   │   │   ├── emq.go
│   │   │   ├── emq_client
│   │   │   ├── emq_cmd
│   │   │   ├── protobuf
│   │   │   ├── topic_publish_handler
│   │   │   ├── topic_router
│   │   │   └── topic_subscribe_handler
│   │   ├── logger
│   │   │   ├── log_util.go
│   │   │   ├── logger.go
│   │   │   └── vlogger.go
│   │   ├── main.go
│   │   ├── middleware
│   │   │   ├── auth_middle.go
│   │   │   └── cors
│   │   ├── model
│   │   │   ├── model_base
│   │   │   ├── threat.go
│   │   │   ├── user.go
│   │   │   ├── vehicle_info.go
│   │   │   └── white_list.go
│   │   ├── response
│   │   │   ├── response.go
│   │   │   ├── response_code.go
│   │   │   └── response_msg.go
│   │   ├── router
│   │   │   ├── router.go
│   │   │   └── v1
│   │   ├── service
│   │   │   └── jwt_service.go
│   │   ├── timing
│   │   │   ├── license_timer.go
│   │   │   ├── timing.go
│   │   │   └── timing_pkg
│   │   └── util
│   │       ├── composite_data_utils.go
│   │       ├── compress_util.go
│   │       ├── data_type_util.go
│   │       ├── encryption_util.go
│   │       ├── file_util.go
│   │       ├── http_client.go
│   │       ├── ip_util.go
│   │       ├── log_util.go
│   │       ├── string_util.go
│   │       ├── time_util.go
│   │       └── vjson.go
│   └── vehicle_script
│       ├── api_script
│       │   ├── auth_api.go
│       │   └── regist_api.go
│       ├── emp_service
│       │   ├── emq_service.go
│       │   ├── emq_tls_conf.go
│       │   └── protobuf
│       ├── emq_script
│       │   ├── push_vehicle_info.go
│       │   └── push_vehicle_threat.go
│       └── tool
│           ├── conf_util.go
│           ├── http_util.go
│           └── string_util.go


│   │   ├── conf  ini库读取配置
│   │   │   ├── conf.go
│   │   │   └── conf_util.go
│   │   ├── db   数据库包
│   │   │   ├── db.go
│   │   │   ├── mysql  gorm封装包
│   │   │   └── redis  gredis封装包
│   │   ├── emq   emq包
│   │   │   ├── emq.go
│   │   │   ├── emq_client emq服务初始化
│   │   │   ├── emq_cmd 发布emq消息封装体
│   │   │   ├── protobuf protobuf文件目录
│   │   │   ├── topic_publish_handler 发布emq消息
│   │   │   ├── topic_router emq router分发
│   │   │   └── topic_subscribe_handler 接收emq消息
│   │   ├── logger logger日志目录
│   │   │   ├── log_util.go
│   │   │   ├── logger.go
│   │   │   └── vlogger.go
│   │   ├── main.go 服务入口
│   │   ├── middleware 中间件jwt检测token，跨域
│   │   │   ├── auth_middle.go
│   │   │   └── cors
│   │   ├── model
│   │   │   ├── model_base model接口
│   │   │   ├── threat.go 
│   │   │   ├── user.go
│   │   │   ├── vehicle_info.go
│   │   │   └── white_list.go
│   │   ├── response  response包
│   │   │   ├── response.go
│   │   │   ├── response_code.go
│   │   │   └── response_msg.go
│   │   ├── router 路由
│   │   │   ├── router.go
│   │   │   └── v1
│   │   ├── service
│   │   │   └── jwt_service.go jwt权限
│   │   └── util 工具类
│   │       ├── composite_data_utils.go
│   │       ├── compress_util.go
│   │       ├── data_type_util.go
│   │       ├── encryption_util.go
│   │       ├── file_util.go
│   │       ├── http_client.go
│   │       ├── ip_util.go
│   │       ├── log_util.go
│   │       ├── string_util.go
│   │       ├── time_util.go
│   │       └── vjson.go
│   └── vehicle_script  项目脚本服务目录
│       ├── api_script  api的脚本
│       │   ├── auth_api.go
│       │   └── regist_api.go
│       ├── emp_service emq服务
│       │   ├── emq_service.go
│       │   ├── emq_tls_conf.go
│       │   └── protobuf
│       ├── emq_script emq服务脚本
│       │   ├── push_vehicle_info.go
│       │   └── push_vehicle_threat.go
│       └── tool 工具类
│           ├── conf_util.go
│           ├── http_util.go
│           └── string_util.go

```
