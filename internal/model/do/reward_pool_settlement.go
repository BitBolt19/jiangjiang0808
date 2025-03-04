// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RewardPoolSettlement is the golang structure of table reward_pool_settlement for DAO operations like Where/Data.
type RewardPoolSettlement struct {
	g.Meta    `orm:"table:reward_pool_settlement, do:true"`
	Id        interface{} //
	PoolId    interface{} // 轮次
	Owners    interface{} // 中奖人列表
	TxHash    interface{} //
	TxTime    *gtime.Time //
	TxEventId interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
