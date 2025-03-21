package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"nez-server/internal/controller/announcement"
	"nez-server/internal/controller/lottery"
	"nez-server/internal/controller/nft"
	"nez-server/internal/controller/staking"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					nft.NewV1(),
					staking.NewV1(),
					lottery.NewV1(),
					announcement.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
