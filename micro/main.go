package main

import (
	"basic"
	"basic/common"
	"basic/config"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/config/source/grpc"
	"log"

	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
	"github.com/opentracing/opentracing-go"
	tracer "plugins/tracer/jaeger"
	"plugins/tracer/opentracing/stdhttp"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func init() {
	plugin.Register(cors.NewPlugin())

	plugin.Register(plugin.NewPlugin(
		plugin.WithName("tracer"),
		plugin.WithHandler(
			stdhttp.TracerWrapper,
		),
	))
}

const name = "API gateway"

func main() {
	// 初始化配置
	initCfg()

	stdhttp.SetSamplingFrequency(100)
	t, io, err := tracer.NewTracer(name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	registryObj := consul.NewRegistry(registryOptions)
	ops := micro.Registry(registryObj)
	cmd.Init(ops)
}
func registryOptions(ops *registry.Options) {
	consulCfg := &common.Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)}
	fmt.Println(ops.Addrs)
}
func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	return
}