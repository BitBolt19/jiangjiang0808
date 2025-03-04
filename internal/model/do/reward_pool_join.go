// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RewardPoolJoin is the golang structure of table reward_pool_join for DAO operations like Where/Data.
type RewardPoolJoin struct {
	g.Meta    `orm:"table:reward_pool_join, do:true"`
	Id        interface{} //
	Account   interface{} // 账户
	PoolId    interface{} // 轮次
	TxHash    interface{} //
	TxTime    *gtime.Time //
	TxEventId interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
