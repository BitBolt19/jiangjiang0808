// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "nez-server/api/lottery/v1"
)

type (
	IApiLottery interface {
		// 获取用户拥有的 nft
		GetLotterySett(ctx context.Context, req *v1.GetLotterySettReq) (res *v1.GetLotterySettRes, err error)
	}
)

var (
	localApiLottery IApiLottery
)

func ApiLottery() IApiLottery {
	if localApiLottery == nil {
		panic("implement not found for interface IApiLottery, forgot register?")
	}
	return localApiLottery
}

func RegisterApiLottery(i IApiLottery) {
	localApiLottery = i
}
