// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RewardPoolSettlement is the golang structure for table reward_pool_settlement.
type RewardPoolSettlement struct {
	Id        uint        `json:"id"        orm:"id"          ` //
	PoolId    uint        `json:"poolId"    orm:"pool_id"     ` // 轮次
	Owners    string      `json:"owners"    orm:"owners"      ` // 中奖人列表
	TxHash    string      `json:"txHash"    orm:"tx_hash"     ` //
	TxTime    *gtime.Time `json:"txTime"    orm:"tx_time"     ` //
	TxEventId int         `json:"txEventId" orm:"tx_event_id" ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at"  ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at"  ` //
}
