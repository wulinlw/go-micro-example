module auth

go 1.12

require (
	basic v0.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.14.0
	github.com/micro/go-plugins v1.4.0
	github.com/opentracing/opentracing-go v1.1.0
	plugins v0.0.0
)

replace (
	basic v0.0.0 => ../basic
	plugins v0.0.0 => ../plugins
)
