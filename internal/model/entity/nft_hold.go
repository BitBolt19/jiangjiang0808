// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NftHold is the golang structure for table nft_hold.
type NftHold struct {
	Id              uint        `json:"id"              orm:"id"               ` //
	Account         string      `json:"account"         orm:"account"          ` // 账户
	ContractAddress string      `json:"contractAddress" orm:"contract_address" ` // NFT合约地址
	TokenId         uint        `json:"tokenId"         orm:"token_id"         ` // NFT tokenId
	TxEventId       int         `json:"txEventId"       orm:"tx_event_id"      ` //
	HoldTime        *gtime.Time `json:"holdTime"        orm:"hold_time"        ` // 持有时间
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       ` //
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       ` //
}
