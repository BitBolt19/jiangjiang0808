package lottery

import (
	"context"
	"nez-server/internal/service"

	"nez-server/api/lottery/v1"
)

func (c *ControllerV1) GetLotterySett(ctx context.Context, req *v1.GetLotterySettReq) (res *v1.GetLotterySettRes, err error) {
	return service.ApiLottery().GetLotterySett(ctx, req)
}
