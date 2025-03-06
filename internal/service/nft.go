// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"nez-server/internal/logic/contract/buy"
	"nez-server/internal/model/entity"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethscanner/ethereum-log-scanner/core/scanner"
)

type (
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
		NewNFTTransfer(ctx context.Context, contractAddress common.Address, transfer *buy.BuyBuyNft, log *scanner.Elog) error
		GetLastTransfer(ctx context.Context, contractAddress string, to string, tokenId uint64) (*entity.NftTransfer, error)
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
