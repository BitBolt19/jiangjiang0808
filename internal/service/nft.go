// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "nez-server/api/Nft/v1"
	"nez-server/internal/model/entity"
	"time"
)

type (
	INFT interface {
		// ConsumeEvents 消费链上事件
		ConsumeEvents(ctx context.Context) error
		GetLastTransfer(ctx context.Context, contractAddress string, to string, tokenId uint64) (*entity.NftTransfer, error)
	}
	INFTHold interface {
		NewTransfer(ctx context.Context, contract string, to string, tokenId uint64, holdTime time.Time) error
		GetAccountHold(ctx context.Context, req *v1.GetNftHoldReq) (*v1.GetNftHoldRes, error)
		GetAccountNFTHoldNum(ctx context.Context, account string) (holdNum int, err error)
	}
)

var (
	localNFT     INFT
	localNFTHold INFTHold
)

func NFT() INFT {
	if localNFT == nil {
		panic("implement not found for interface INFT, forgot register?")
	}
	return localNFT
}

func RegisterNFT(i INFT) {
	localNFT = i
}

func NFTHold() INFTHold {
	if localNFTHold == nil {
		panic("implement not found for interface INFTHold, forgot register?")
	}
	return localNFTHold
}

func RegisterNFTHold(i INFTHold) {
	localNFTHold = i
}
