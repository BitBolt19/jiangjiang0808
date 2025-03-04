package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"nez-server/internal/apiEntity"
)

type GetNftHoldReq struct {
	g.Meta  `path:"/nft/hold" tags:"nft" method:"get" summary:"nft持有列表"`
	Account string `json:"account" dc:"用户地址"`
}

type GetNftHoldRes struct {
	List []*apiEntity.NFTHoldInfo
}
