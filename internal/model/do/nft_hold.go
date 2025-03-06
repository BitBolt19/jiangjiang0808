// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NftHold is the golang structure of table nft_hold for DAO operations like Where/Data.
type NftHold struct {
	g.Meta          `orm:"table:nft_hold, do:true"`
	Id              interface{} //
	Account         interface{} // 账户
	ContractAddress interface{} // NFT合约地址
	TokenId         interface{} // NFT tokenId
	TxEventId       interface{} //
	HoldTime        *gtime.Time // 持有时间
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
