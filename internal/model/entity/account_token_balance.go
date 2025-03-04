// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AccountTokenBalance is the golang structure for table account_token_balance.
type AccountTokenBalance struct {
	Id        uint        `json:"id"        orm:"id"         ` //
	Account   string      `json:"account"   orm:"account"    ` // 账户
	Token     string      `json:"token"     orm:"token"      ` // 代币
	Balance   float64     `json:"balance"   orm:"balance"    ` // 余额
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
}
