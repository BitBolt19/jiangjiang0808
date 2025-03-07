package apiStaking

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "nez-server/api/staking/v1"
	"nez-server/internal/apiEntity"
	"nez-server/internal/dao"
	"nez-server/internal/service"
)

type (
	sApiStaking struct{}
)

func init() {
	service.RegisterApiStaking(NewApiStaking())
}

func NewApiStaking() service.IApiStaking {
	return &sApiStaking{}
}

// 获取用户的NFT质押信息
func (s *sApiStaking) GetUserStaking(ctx context.Context, req *v1.GetUserStakingReq) (res *v1.GetUserStakingRes, err error) {
	account := req.Account
	result, err := dao.NftStaking.Ctx(ctx).Where("account = ?", account).Fields("account,tx_time").All()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	var list []*apiEntity.UserStakingInfo
	if err = result.Structs(&list); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &v1.GetUserStakingRes{
		List: list,
	}, nil
}
