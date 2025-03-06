package apiEntity

type NFTHoldInfo struct {
	Account         string `json:"account" dc:"用户地址"`
	ContractAddress string `json:"token" dc:"合约地址"`
	Count           uint   `json:"count" dc:"nft张数"`
}
