// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "nez-server/api/staking/v1"
)

type (
	IApiStaking interface {
		// 获取用户的NFT质押信息
		GetUserStaking(ctx context.Context, req *v1.GetUserStakingReq) (res *v1.GetUserStakingRes, err error)
	}
)

var (
	localApiStaking IApiStaking
)

func ApiStaking() IApiStaking {
	if localApiStaking == nil {
		panic("implement not found for interface IApiStaking, forgot register?")
	}
	return localApiStaking
}

func RegisterApiStaking(i IApiStaking) {
	localApiStaking = i
}
