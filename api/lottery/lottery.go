// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package lottery

import (
	"context"

	"nez-server/api/lottery/v1"
)

type ILotteryV1 interface {
	GetLotterySett(ctx context.Context, req *v1.GetLotterySettReq) (res *v1.GetLotterySettRes, err error)
}
