// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AccountTokenBalanceLog is the golang structure of table account_token_balance_log for DAO operations like Where/Data.
type AccountTokenBalanceLog struct {
	g.Meta     `orm:"table:account_token_balance_log, do:true"`
	Id         interface{} //
	OrderId    interface{} // 关联订单ID
	Account    interface{} // 账户
	Token      interface{} // 代币
	Amount     interface{} // 数量
	OldBalance interface{} // 旧余额
	NewBalance interface{} // 新余额
	Type       interface{} // 流水类型
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
