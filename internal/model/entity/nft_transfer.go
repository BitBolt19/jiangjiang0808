// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NftTransfer is the golang structure for table nft_transfer.
type NftTransfer struct {
	Id              uint        `json:"id"              orm:"id"               ` //
	Account         string      `json:"account"         orm:"account"          ` // 账户
	ContractAddress string      `json:"contractAddress" orm:"contract_address" ` // NFT合约地址
	FromAddress     string      `json:"fromAddress"     orm:"from_address"     ` // 发送方
	ToAddress       string      `json:"toAddress"       orm:"to_address"       ` // 接收方
	Amount          int64       `json:"amount"          orm:"amount"           ` // 数量
	TokenId         uint        `json:"tokenId"         orm:"token_id"         ` // NFT tokenId
	TxHash          string      `json:"txHash"          orm:"tx_hash"          ` // 交易hash
	TxEventId       uint        `json:"txEventId"       orm:"tx_event_id"      ` // 交替事件id
	TxTime          *gtime.Time `json:"txTime"          orm:"tx_time"          ` // 交易时间
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       ` //
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       ` //
}
