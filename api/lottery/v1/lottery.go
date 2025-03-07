package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"nez-server/internal/apiEntity"
)

type GetLotterySettReq struct {
	g.Meta `path:"/lottery/rpSett" tags:"抽奖" method:"get" summary:"中奖用户列表"`
	PoolId uint `json:"pool_id" dc:"中奖轮次"`
}

type GetLotterySettRes struct {
	List []*apiEntity.LotterySettInfo
}
