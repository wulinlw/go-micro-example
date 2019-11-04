module user-srv

go 1.12

require (
	basic v0.0.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.14.0
	github.com/micro/go-plugins v1.4.0
	github.com/opentracing/opentracing-go v1.1.0
	go.uber.org/zap v1.10.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	plugins v0.0.0
)

replace (
	basic v0.0.0 => ../basic
	plugins v0.0.0 => ../plugins
)
