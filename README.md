# vehicle

仓促完成，后续在更新！！！

用到的技术
gin、goredis、gorm、restfulapi、protobuf、emq、jwt、gomod

项目的总体构建流程如下:
![image](https://github.com/iyongfei/vehicle/blob/master/assets/outer_line.png)



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
│   │   │   ├── mysql gorm封装
│   │   │   └── redis goredis封装
│   │   ├── docs   swagger测试
│   │   │   ├── docs.go
│   │   │   ├── swagger.json
│   │   │   └── swagger.yaml
│   │   ├── emq emq包
│   │   │   ├── emq.go
│   │   │   ├── emq_client
│   │   │   ├── emq_cmd
│   │   │   ├── protobuf
│   │   │   ├── topic_publish_handler
│   │   │   ├── topic_router
│   │   │   └── topic_subscribe_handler
│   │   ├── logger 日志包
│   │   │   ├── log_util.go
│   │   │   ├── logger.go
│   │   │   └── vlogger.go
│   │   ├── main.go
│   │   ├── middleware 中间件
│   │   │   ├── auth_middle.go
│   │   │   └── cors
│   │   ├── model model包
│   │   │   ├── model_base
│   │   │   ├── threat.go
│   │   │   ├── user.go
│   │   │   ├── vehicle_info.go
│   │   │   └── white_list.go
│   │   ├── response 返回格式封装
│   │   │   ├── response.go
│   │   │   ├── response_code.go
│   │   │   └── response_msg.go
│   │   ├── router 路由版本
│   │   │   ├── router.go
│   │   │   └── v1
│   │   ├── service jwt权限
│   │   │   └── jwt_service.go
│   │   ├── timing  倒计时
│   │   │   ├── license_timer.go
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
│   └── vehicle_script 项目的测试包
│       ├── api_script api接口测试脚本
│       │   ├── auth_api.go
│       │   └── regist_api.go
│       ├── emp_service 
│       │   ├── emq_service.go
│       │   ├── emq_tls_conf.go
│       │   └── protobuf
│       ├── emq_script emq测试脚本
│       │   ├── push_vehicle_info.go
│       │   └── push_vehicle_threat.go
│       └── tool
│           ├── conf_util.go
│           ├── http_util.go
│           └── string_util.go


```
