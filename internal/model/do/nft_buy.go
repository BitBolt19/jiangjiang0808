// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NftBuy is the golang structure of table nft_buy for DAO operations like Where/Data.
type NftBuy struct {
	g.Meta          `orm:"table:nft_buy, do:true"`
	Id              interface{} //
	Account         interface{} // 账户
	ContractAddress interface{} // NFT合约地址
	TokenId         interface{} // NFT tokenId
	TxHash          interface{} // 交易hash
	TxTime          *gtime.Time // 购买盲盒时间
	TxEventId       interface{} // 交易事件id
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
