module user-web

go 1.12

require (
	auth v0.0.0
	basic v0.0.0
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.14.0
	github.com/micro/go-plugins v1.4.0
	github.com/opentracing/opentracing-go v1.1.0
	plugins v0.0.0
	user-srv v0.0.0
)

replace (
	auth v0.0.0 => ../auth
	basic v0.0.0 => ../basic
	plugins v0.0.0 => ../plugins
	user-srv v0.0.0 => ../user-srv
)
