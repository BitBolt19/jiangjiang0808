// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NftBuy is the golang structure for table nft_buy.
type NftBuy struct {
	Id              uint        `json:"id"              orm:"id"               ` //
	Account         string      `json:"account"         orm:"account"          ` // 账户
	ContractAddress string      `json:"contractAddress" orm:"contract_address" ` // NFT合约地址
	TokenId         uint        `json:"tokenId"         orm:"token_id"         ` // NFT tokenId
	TxHash          string      `json:"txHash"          orm:"tx_hash"          ` // 交易hash
	TxTime          *gtime.Time `json:"txTime"          orm:"tx_time"          ` // 购买盲盒时间
	TxEventId       uint        `json:"txEventId"       orm:"tx_event_id"      ` // 交易事件id
	Status          int         `json:"status"          orm:"status"           ` // 盲盒状态
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       ` //
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       ` //
}
