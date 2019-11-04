# go-micro-example
go-micro的项目例子，参考https://github.com/micro-in-cn/tutorials/blob/master/microservice-in-micro/README.md  
本地测试使用go 1.13 ，使用module管理依赖    

公共组件目录basic及插件目录plugins，需要预先设置好go mod,注意在go.mod引用处replace成本地相对路径  
启动docker-compose up  
启动config-grpc-srv 配置中心，必须先启动，所有项目启动时都从这里获取配置  
启动auth  
启动user-srv  
启动user-web  
调试时可以启动micro --registry=consul  web  
可在micro web中看到mu.micro.book.srv.user  
测试调用内容{"userID":"10001","userName":"micro","userPwd":"1234"}  
默认情况下api模式  
micro.exe --registry=consul --api_namespace=mu.micro.book.web  api --handler=web --address=127.0.0.1:8080  

全链路跟踪测试时，使用micro目录的自定义版本micro
cusmicro.exe  --registry=consul --api_namespace=mu.micro.book.web  api --handler=web 
part3中有postman调用代码，导入即可测试  
这里的代码与参考教程第7章有差异，教程中使用的服务发现是etcd，我们本地测试使用的consul，  
在micro 1.15版本中，默认registry不再包含consul，而是移到plugins中了，  
所以必须手动在代码中指定registry使用consul才可正常使用，否则启动参数指定--registry=consul时程序不会报错，而是显示micro的默认提示信息  

consul ui  
http://localhost:8500/  

熔断面板  
浏览器打开 http://localhost:8081/hystrix.stream  
输入 http://{ip}:81/hystrix.stream , 点击最下方按钮。此处ip为本机ip，因为hystrix-dashboard是容器启动的，无法直接访问本机127.0.0.1  

jaeger链路跟踪  
http://localhost:16686/  