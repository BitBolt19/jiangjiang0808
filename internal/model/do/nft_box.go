// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NftBox is the golang structure of table nft_box for DAO operations like Where/Data.
type NftBox struct {
	g.Meta          `orm:"table:nft_box, do:true"`
	Id              interface{} //
	Account         interface{} // 账户
	ContractAddress interface{} // NFT合约地址
	TokenId         interface{} // NFT tokenId
	Amount          interface{} // 数量
	TxHash          interface{} // 交易hash
	TxTime          *gtime.Time // 开盲盒时间
	TxEventId       interface{} // 交易事件id
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
