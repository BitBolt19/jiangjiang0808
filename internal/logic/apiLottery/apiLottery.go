package apiLottery

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "nez-server/api/lottery/v1"
	"nez-server/internal/apiEntity"
	"nez-server/internal/dao"
	"nez-server/internal/service"
)

type (
	sApiLottery struct{}
)

func init() {
	service.RegisterApiLottery(NewApiLottery())
}

func NewApiLottery() service.IApiLottery {
	return &sApiLottery{}
}

// 获取用户拥有的 nft
func (s *sApiLottery) GetLotterySett(ctx context.Context, req *v1.GetLotterySettReq) (res *v1.GetLotterySettRes, err error) {
	poolId := req.PoolId
	result := dao.RewardPoolSettlement.Ctx(ctx).Fields("pool_id,owners,tx_time")
	if poolId != 0 {
		result = result.Where("pool_id = ?", poolId)
	}
	resSett, err := result.All()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	var list []*apiEntity.LotterySettInfo
	if err = resSett.Structs(&list); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &v1.GetLotterySettRes{
		List: list,
	}, nil
}
