// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "nez-server/api/nft/v1"
)

type (
	IApiNft interface {
		// 获取用户拥有的 nft
		GetUserNft(ctx context.Context, req *v1.GetNftHoldReq) (res *v1.GetNftHoldRes, err error)
	}
)

var (
	localApiNft IApiNft
)

func ApiNft() IApiNft {
	if localApiNft == nil {
		panic("implement not found for interface IApiNft, forgot register?")
	}
	return localApiNft
}

func RegisterApiNft(i IApiNft) {
	localApiNft = i
}
