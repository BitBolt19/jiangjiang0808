// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"nez-server/internal/logic/contract/buy"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethscanner/ethereum-log-scanner/core/scanner"
)

type (
	INFTbox interface {
		// ConsumeEvents 消费链上事件
		ConsumeEvents(ctx context.Context) error
		// 当盲盒打开后，购买 状态 要更改为已开启
		UpdateStatus(ctx context.Context, event *buy.BuyOpenBlindBox, eventId uint) (err error)
	}
	INFTBuy interface {
		// ConsumeEvents 消费链上事件
		ConsumeEvents(ctx context.Context) error
	}
	INFTHold interface {
		NewHold(ctx context.Context, contract string, to string, tokenId uint64, holdTime time.Time) error
		GetAccountNFTHoldNum(ctx context.Context, account string) (holdNum int, err error)
	}
	INFTStaking interface {
		// ConsumeEvents 消费链上事件
		ConsumeEvents(ctx context.Context) error
		GetStakingInfo(ctx context.Context, account string) (map[string]uint64, error)
	}
	INFT interface {
		// ConsumeEvents 消费链上事件
		ConsumeEvents(ctx context.Context) error
		NewTransfer(ctx context.Context, contract common.Address, from common.Address, to common.Address, amount int64, log *scanner.Elog) error
	}
)

var (
	localNFTbox     INFTbox
	localNFTBuy     INFTBuy
	localNFTHold    INFTHold
	localNFTStaking INFTStaking
	localNFT        INFT
)

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

func NFT() INFT {
	if localNFT == nil {
		panic("implement not found for interface INFT, forgot register?")
	}
	return localNFT
}

func RegisterNFT(i INFT) {
	localNFT = i
}
