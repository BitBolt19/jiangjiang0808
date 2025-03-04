// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RewardPoolClaim is the golang structure for table reward_pool_claim.
type RewardPoolClaim struct {
	Id        uint        `json:"id"        orm:"id"          ` //
	Account   string      `json:"account"   orm:"account"     ` // 账户
	PoolId    uint        `json:"poolId"    orm:"pool_id"     ` // 轮次
	Token     string      `json:"token"     orm:"token"       ` // 代币合约地址
	Amount    float64     `json:"amount"    orm:"amount"      ` // 数量
	TxHash    string      `json:"txHash"    orm:"tx_hash"     ` //
	TxTime    *gtime.Time `json:"txTime"    orm:"tx_time"     ` //
	TxEventId uint        `json:"txEventId" orm:"tx_event_id" ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at"  ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at"  ` //
}
