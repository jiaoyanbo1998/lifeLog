package main

import (
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	lifelogv1 "lifelog-grpc/api/proto/gen/api/proto/lifelog/v1"
	"strconv"
)

func main() {
	// 1.解析配置文件
	type config struct {
		Port int      `yaml:"port"`
		Addr []string `yaml:"addr"`
	}
	cfg := config{
		Port: 8072,
		Addr: []string{"localhost:12379"},
	}
	err := viper.UnmarshalKey("lifeLogEtcd", &cfg)
	if err != nil {
		panic(err)
	}
	// 2.grpc服务器连接的配置（go-zero框架）
	serverConf := zrpc.RpcServerConf{
		// lifelog服务监听的端口
		ListenOn: ":" + strconv.Itoa(cfg.Port),
		// 服务注册
		Etcd: discov.EtcdConf{
			Hosts: cfg.Addr,          // etcd地址
			Key:   "service/lifelog", // lifelog服务的唯一标识
		},
	}
	// 3.创建lifelog服务的grpc服务器
	lifeLogServiceGRPCService := InitLifeLogServiceGRPCService()
	// 4.创建grpc服务器（go-zero）
	server := zrpc.MustNewServer(serverConf, func(server *grpc.Server) {
		// 5.将lifelog服务注册到grpc服务器中
		lifelogv1.RegisterLifeLogServiceServer(server, lifeLogServiceGRPCService)
	})
	// 6.关闭grpc服务器
	defer server.Stop()
	// 7.启动grpc服务器
	logx.Info("正在启动lifelogService，服务端口：", cfg.Port, "......")
	server.Start()
}
