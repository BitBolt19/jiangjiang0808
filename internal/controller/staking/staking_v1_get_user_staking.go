package staking

import (
	"context"
	"nez-server/internal/service"

	"nez-server/api/staking/v1"
)

func (c *ControllerV1) GetUserStaking(ctx context.Context, req *v1.GetUserStakingReq) (res *v1.GetUserStakingRes, err error) {
	return service.ApiStaking().GetUserStaking(ctx, req)
}
