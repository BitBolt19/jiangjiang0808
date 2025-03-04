// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HdContractEvent is the golang structure for table hd_contract_event.
type HdContractEvent struct {
	Id              uint64      `json:"id"              orm:"id"               ` // id
	ContractName    string      `json:"contractName"    orm:"contract_name"    ` // 合约名
	ContractAddress string      `json:"contractAddress" orm:"contract_address" ` // 合约地址
	TxHash          string      `json:"txHash"          orm:"tx_hash"          ` // 交易哈希
	EventHash       string      `json:"eventHash"       orm:"event_hash"       ` // 事件名
	EventId         int64       `json:"eventId"         orm:"event_id"         ` // 事件id
	BlockNumber     int64       `json:"blockNumber"     orm:"block_number"     ` // 区块编号
	BlockHash       string      `json:"blockHash"       orm:"block_hash"       ` // 交易哈希
	EventTopics     []byte      `json:"eventTopics"     orm:"event_topics"     ` // 事件头
	EventData       []byte      `json:"eventData"       orm:"event_data"       ` // event数据
	State           int         `json:"state"           orm:"state"            ` // 0:待处理 10:已处理
	CheckState      int         `json:"checkState"      orm:"check_state"      ` // 链上状态: 0:待处理 10:已确认 20:确认异常
	CheckedBlock    uint64      `json:"checkedBlock"    orm:"checked_block"    ` // 已确认区块
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       ` // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       ` // 更新时间
}
