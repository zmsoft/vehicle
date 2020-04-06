# vehicle

用到的技术
gin、goredis、gorm、restfulapi、protobuf、emq、jwt、gomod

项目的总体构建流程如下:
![image](https://github.com/iyongfei/vehicle/blob/master/assets/outer_line.png)

.
├── README.md
├── assets
│   ├── outer_line.png
│   └── vehicle.sql
├── conf.ini
├── conf.txt
├── go.mod
├── go.sum
├── src
│   ├── vehicle
│   │   ├── api_server
│   │   │   ├── auth.go
│   │   │   ├── threat.go
│   │   │   ├── vehicle.go
│   │   │   └── white_list.go
│   │   ├── conf
│   │   │   ├── conf.go
│   │   │   └── conf_util.go
│   │   ├── db
│   │   │   ├── db.go
│   │   │   ├── mysql
│   │   │   │   ├── add.go
│   │   │   │   ├── dele.go
│   │   │   │   ├── gorm.go
│   │   │   │   ├── query.go
│   │   │   │   └── update.go
│   │   │   └── redis
│   │   │       ├── redis_db.go
│   │   │       └── redis_util.go
│   │   ├── emq
│   │   │   ├── emq.go
│   │   │   ├── emq_client
│   │   │   │   ├── conn_lost_handler.go
│   │   │   │   ├── emq_client.go
│   │   │   │   └── tls_config.go
│   │   │   ├── emq_cmd
│   │   │   │   ├── emq_cmd_impl.go
│   │   │   │   └── vehicle_cmd.go
│   │   │   ├── protobuf
│   │   │   │   ├── gwresult.pb.go
│   │   │   │   ├── gwresult.proto
│   │   │   │   ├── protobuf.txt
│   │   │   │   ├── scommand.pb.go
│   │   │   │   └── scommand.proto
│   │   │   ├── topic_publish_handler
│   │   │   │   ├── publish_msg.go
│   │   │   │   └── publish_service.go
│   │   │   ├── topic_router
│   │   │   │   ├── topic_router.go
│   │   │   │   ├── topic_router_data.go
│   │   │   │   └── topic_router_data_pool.go
│   │   │   └── topic_subscribe_handler
│   │   │       ├── action_handler_dispatch.go
│   │   │       ├── vehicle_info_handle.go
│   │   │       ├── vehicle_online.go
│   │   │       ├── vehicle_strategy_handle.go
│   │   │       └── vehicle_threat_handle.go
│   │   ├── logger
│   │   │   ├── log_util.go
│   │   │   ├── logger.go
│   │   │   └── vlogger.go
│   │   ├── main.go
│   │   ├── middleware
│   │   │   ├── auth_middle.go
│   │   │   └── cors
│   │   │       ├── cors.go
│   │   │       ├── cors_conf.go
│   │   │       └── cors_util.go
│   │   ├── model
│   │   │   ├── model_base
│   │   │   │   └── model_base.go
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
│   │   │       └── v1.go
│   │   ├── service
│   │   │   └── jwt_service.go
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
│       │       ├── gwresult.pb.go
│       │       ├── gwresult.proto
│       │       ├── protobuf.txt
│       │       ├── scommand.pb.go
│       │       └── scommand.proto
│       ├── emq_script
│       │   ├── push_vehicle_info.go
│       │   └── push_vehicle_threat.go
│       └── tool
│           ├── conf_util.go
│           ├── http_util.go
│           └── string_util.go
