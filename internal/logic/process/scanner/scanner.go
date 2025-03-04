package scanner

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"nez-server/internal/service"
	"sync"
)

func Scanner() {
	ctx := gctx.GetInitCtx()
	_, err := gcron.AddSingleton(ctx, "*/10 * * * * *", func(ctx context.Context) {
		var wg sync.WaitGroup
		wg.Add(3)
		go func() {
			defer wg.Done()
			err := service.NFTStaking().ConsumeEvents(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "failed to consume account bind events, error:%+v", err)
			}
		}()
		go func() {
			defer wg.Done()
			err := service.NFT().ConsumeEvents(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "failed to consume NFT events, error:%+v", err)
			}
		}()
		go func() {
			defer wg.Done()
			err := service.NFTBuy().ConsumeEvents(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "failed to consume NFT Buy events, error:%+v", err)
			}
		}()
		wg.Wait()
	})
	if err != nil {
		g.Log().Error(ctx, err)
	}
	//select {}
}
