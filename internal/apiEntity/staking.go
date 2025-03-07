package apiEntity

type UserStakingInfo struct {
	Account string        `json:"account" dc:"用户地址"`
	TxTime  TimestampDate `json:"tx_time" dc:"质押时间"`
}
