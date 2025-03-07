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
		wg.Add(6)
		// nft 质押
		go func() {
			defer wg.Done()
			err := service.NFTStaking().ConsumeEvents(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "failed to consume account bind events, error:%+v", err)
			}
		}()
		// nft Transfer
		go func() {
			defer wg.Done()
			err := service.NFT().ConsumeEvents(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "failed to consume NFT events, error:%+v", err)
			}
		}()
		// nft 购买
		go func() {
			defer wg.Done()
			err := service.NFTBuy().ConsumeEvents(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "failed to consume NFT Buy events, error:%+v", err)
			}
		}()
		// nft 开盲盒
		go func() {
			defer wg.Done()
			err := service.NFTbox().ConsumeEvents(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "failed to consume NFT Box events, error:%+v", err)
			}
		}()
		//  参与抽奖
		go func() {
			defer wg.Done()
			err := service.RpJoin().ConsumeEvents(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "failed to consume reward join events, error:%+v", err)
			}
		}()
		//  开奖
		go func() {
			defer wg.Done()
			err := service.RpSett().ConsumeEvents(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "failed to consume Reward settement events, error:%+v", err)
			}
		}()
		wg.Wait()
	})
	if err != nil {
		g.Log().Error(ctx, err)
	}
	//select {}
}
