module plugins

go 1.12

require (
	basic v0.0.0
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/google/uuid v1.1.1
	github.com/gorilla/sessions v1.2.0
	github.com/micro/go-micro v1.14.0
	github.com/micro/go-plugins v1.4.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.19.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
)

replace basic v0.0.0 => ../basic
