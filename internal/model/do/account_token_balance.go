// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AccountTokenBalance is the golang structure of table account_token_balance for DAO operations like Where/Data.
type AccountTokenBalance struct {
	g.Meta    `orm:"table:account_token_balance, do:true"`
	Id        interface{} //
	Account   interface{} // 账户
	Token     interface{} // 代币
	Balance   interface{} // 余额
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
