// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NftTransfer is the golang structure of table nft_transfer for DAO operations like Where/Data.
type NftTransfer struct {
	g.Meta          `orm:"table:nft_transfer, do:true"`
	Id              interface{} //
	Account         interface{} // 账户
	ContractAddress interface{} // NFT合约地址
	FromAddress     interface{} // 发送方
	ToAddress       interface{} // 接收方
	Amount          interface{} // 数量
	TokenId         interface{} // NFT tokenId
	TxHash          interface{} // 交易hash
	TxEventId       interface{} // 交替事件id
	TxTime          *gtime.Time // 交易时间
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
