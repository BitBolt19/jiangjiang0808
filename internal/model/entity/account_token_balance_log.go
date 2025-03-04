// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AccountTokenBalanceLog is the golang structure for table account_token_balance_log.
type AccountTokenBalanceLog struct {
	Id         uint        `json:"id"         orm:"id"          ` //
	OrderId    uint        `json:"orderId"    orm:"order_id"    ` // 关联订单ID
	Account    string      `json:"account"    orm:"account"     ` // 账户
	Token      string      `json:"token"      orm:"token"       ` // 代币
	Amount     float64     `json:"amount"     orm:"amount"      ` // 数量
	OldBalance float64     `json:"oldBalance" orm:"old_balance" ` // 旧余额
	NewBalance float64     `json:"newBalance" orm:"new_balance" ` // 新余额
	Type       uint        `json:"type"       orm:"type"        ` // 流水类型
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` //
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  ` //
}
