// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RewardPoolClaim is the golang structure of table reward_pool_claim for DAO operations like Where/Data.
type RewardPoolClaim struct {
	g.Meta    `orm:"table:reward_pool_claim, do:true"`
	Id        interface{} //
	Account   interface{} // 账户
	PoolId    interface{} // 轮次
	Token     interface{} // 代币合约地址
	Amount    interface{} // 数量
	TxHash    interface{} //
	TxTime    *gtime.Time //
	TxEventId interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
