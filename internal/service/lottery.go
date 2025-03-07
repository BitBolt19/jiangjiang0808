// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IRpClaim interface {
		// ConsumeEvents 消费链上事件
		ConsumeEvents(ctx context.Context) error
		GetBuyInfo(ctx context.Context, account string) (map[string]uint64, error)
	}
	IRpJoin interface {
		// ConsumeEvents 消费链上事件
		ConsumeEvents(ctx context.Context) error
		GetBuyInfo(ctx context.Context, account string) (map[string]uint64, error)
	}
	IRpSett interface {
		// ConsumeEvents 消费链上事件
		ConsumeEvents(ctx context.Context) error
		GetBuyInfo(ctx context.Context, account string) (map[string]uint64, error)
	}
)

var (
	localRpClaim IRpClaim
	localRpJoin  IRpJoin
	localRpSett  IRpSett
)

func RpClaim() IRpClaim {
	if localRpClaim == nil {
		panic("implement not found for interface IRpClaim, forgot register?")
	}
	return localRpClaim
}

func RegisterRpClaim(i IRpClaim) {
	localRpClaim = i
}

func RpJoin() IRpJoin {
	if localRpJoin == nil {
		panic("implement not found for interface IRpJoin, forgot register?")
	}
	return localRpJoin
}

func RegisterRpJoin(i IRpJoin) {
	localRpJoin = i
}

func RpSett() IRpSett {
	if localRpSett == nil {
		panic("implement not found for interface IRpSett, forgot register?")
	}
	return localRpSett
}

func RegisterRpSett(i IRpSett) {
	localRpSett = i
}
