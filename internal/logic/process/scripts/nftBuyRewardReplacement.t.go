package scripts

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	_ "nft-staking/internal/logic"
	"nft-staking/internal/service"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

// TestNFTBuyRewardReplacement 手动发放NFT购买奖励
func TestNFTBuyRewardReplacement(t *testing.T) {
	ctx := context.Background()
	err := service.NFTBuy().RunReward(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}
