// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package nft

import (
	"context"

	"nez-server/api/nft/v1"
)

type INftV1 interface {
	GetNftHold(ctx context.Context, req *v1.GetNftHoldReq) (res *v1.GetNftHoldRes, err error)
}
