package nft

import (
	"context"
	"nez-server/internal/service"

	"nez-server/api/nft/v1"
)

func (c *ControllerV1) GetNftHold(ctx context.Context, req *v1.GetNftHoldReq) (res *v1.GetNftHoldRes, err error) {
	return service.ApiNft().GetUserNft(ctx, req)
}
