package main

import (
	"flag"
	"fmt"

	"github.com/CoderSamYhc/learn-go-zero/apps/order/rpc/internal/config"
	"github.com/CoderSamYhc/learn-go-zero/apps/order/rpc/internal/server"
	"github.com/CoderSamYhc/learn-go-zero/apps/order/rpc/internal/svc"
	"github.com/CoderSamYhc/learn-go-zero/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var orderConfigFile = flag.String("f", "etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*orderConfigFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewOrderServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		order.RegisterOrderServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
