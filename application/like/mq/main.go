package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/service"
	"inquora/application/like/mq/internal/logic"

	"inquora/application/like/mq/internal/config"
	"inquora/application/like/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/like.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range logic.Consumers(ctx, svcCtx) {
		serviceGroup.Add(mq)
	}
	serviceGroup.Start()
}
