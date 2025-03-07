package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"nez-server/internal/apiEntity"
)

type GetUserStakingReq struct {
	g.Meta  `path:"/staking/NftStaking" tags:"质押" method:"get" summary:"用户质押列表"`
	Account string `json:"account" dc:"用户地址"`
}

type GetUserStakingRes struct {
	List []*apiEntity.UserStakingInfo
}
