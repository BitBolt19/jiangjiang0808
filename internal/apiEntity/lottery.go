package apiEntity

type LotterySettInfo struct {
	PoolId uint64        `json:"pool_id" dc:"中奖轮次"`
	Owners string        `json:"owners" dc:"中奖用户"`
	TxTime TimestampDate `json:"tx_time" dc:"中奖时间"`
}
