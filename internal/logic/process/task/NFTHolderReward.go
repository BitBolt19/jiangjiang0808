package task

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"nft-staking/internal/service"
	"sync"
)

func RunReward() {
	ctx := gctx.GetInitCtx()
	_, err := gcron.AddSingleton(ctx, "*/1 * * * * *", func(ctx context.Context) {
		var wg sync.WaitGroup
		wg.Add(1)
		//发放NFT持有奖励（GP）
		go func() {
			defer wg.Done()
			err := service.NFTHolderReward().RunReward(ctx)
			if err != nil {
				g.Log().Error(ctx, "failed to run NFT holder reward", err)
			}
		}()
		wg.Wait()
	})
	if err != nil {
		g.Log().Error(ctx, err)
	}
	//select {}
}
