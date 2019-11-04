module cusmicro

go 1.12

require (
	basic v0.0.0
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/micro/go-micro v1.15.1
	github.com/micro/go-plugins v1.4.0
	github.com/micro/micro v1.15.1
	github.com/opentracing/opentracing-go v1.1.0
	plugins v0.0.0
)

replace (
	basic v0.0.0 => ../basic
	plugins v0.0.0 => ../plugins
)
