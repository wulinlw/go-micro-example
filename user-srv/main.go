package main

import (
	"fmt"
	"time"

	"basic"
	"basic/common"
	"basic/config"
	tracer "plugins/tracer/jaeger"
	"user-srv/handler"
	"user-srv/model"
	s "user-srv/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	//"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/grpc"
	ocplugin "github.com/micro/go-plugins/wrapper/trace/opentracing"
	openTrace "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	z "plugins/zap"
)

var (
	appName = "user_srv"
	cfg     = &userCfg{}
	log = z.GetLogger()
)

type userCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置、数据库等信息
	initCfg()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	t, io, err := tracer.NewTracer(cfg.Name, "localhost:6831")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer io.Close()
	openTrace.SetGlobalTracer(t)
	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(micReg),
		micro.Version("latest"),
		micro.WrapHandler(ocplugin.NewHandlerWrapper(t)),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := &common.Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)}
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)
	//config.WithSource  是basic.config中函数，返回的是一个方法
	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	//log.Logf("[initCfg] 配置，cfg：%v", cfg)
	log.Info("[initCfg] 配置", zap.Any("cfg", cfg))
	return
}
