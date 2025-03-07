// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package staking

import (
	"context"

	"nez-server/api/staking/v1"
)

type IStakingV1 interface {
	GetUserStaking(ctx context.Context, req *v1.GetUserStakingReq) (res *v1.GetUserStakingRes, err error)
}
