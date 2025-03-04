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
	INFTbox interface {
		// ConsumeEvents 消费链上事件
		ConsumeEvents(ctx context.Context) error
		GetBuyInfo(ctx context.Context, account string) (map[string]uint64, error)
	}
	INFTBuy interface {
		// ConsumeEvents 消费链上事件
		ConsumeEvents(ctx context.Context) error
		GetBuyInfo(ctx context.Context, account string) (map[string]uint64, error)
	}
	INFTHold interface {
		NewTransfer(ctx context.Context, contract string, to string, tokenId uint64, holdTime time.Time) error
		GetAccountHold(ctx context.Context, req *v1.GetNftHoldReq) (*v1.GetNftHoldRes, error)
		GetAccountNFTHoldNum(ctx context.Context, account string) (holdNum int, err error)
	}
	INFTStaking interface {
		// ConsumeEvents 消费链上事件
		ConsumeEvents(ctx context.Context) error
		GetStakingInfo(ctx context.Context, account string) (map[string]uint64, error)
	}
)

var (
	localNFT        INFT
	localNFTbox     INFTbox
	localNFTBuy     INFTBuy
	localNFTHold    INFTHold
	localNFTStaking INFTStaking
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

func NFTbox() INFTbox {
	if localNFTbox == nil {
		panic("implement not found for interface INFTbox, forgot register?")
	}
	return localNFTbox
}

func RegisterNFTbox(i INFTbox) {
	localNFTbox = i
}

func NFTBuy() INFTBuy {
	if localNFTBuy == nil {
		panic("implement not found for interface INFTBuy, forgot register?")
	}
	return localNFTBuy
}

func RegisterNFTBuy(i INFTBuy) {
	localNFTBuy = i
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

func NFTStaking() INFTStaking {
	if localNFTStaking == nil {
		panic("implement not found for interface INFTStaking, forgot register?")
	}
	return localNFTStaking
}

func RegisterNFTStaking(i INFTStaking) {
	localNFTStaking = i
}
