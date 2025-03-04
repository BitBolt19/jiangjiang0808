// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NftStakingClaim is the golang structure of table nft_staking_claim for DAO operations like Where/Data.
type NftStakingClaim struct {
	g.Meta          `orm:"table:nft_staking_claim, do:true"`
	Id              interface{} //
	Account         interface{} // 账户
	ContractAddress interface{} // NFT合约地址
	TokenId         interface{} // NFT tokenId
	Amount          interface{} // 领取数量
	TxHash          interface{} //
	TxTime          *gtime.Time //
	TxEventId       interface{} //
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
