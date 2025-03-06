package apiNft

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "nez-server/api/nft/v1"
	"nez-server/internal/apiEntity"
	"nez-server/internal/dao"
	"nez-server/internal/service"
)

type (
	sApiNft struct{}
)

func init() {
	service.RegisterApiNft(NewApiNft())
}

func NewApiNft() service.IApiNft {
	return &sApiNft{}
}

// 获取用户拥有的 nft
func (s *sApiNft) GetUserNft(ctx context.Context, req *v1.GetNftHoldReq) (res *v1.GetNftHoldRes, err error) {
	account := req.Account
	result, err := dao.NftBox.Ctx(ctx).Where("account = ?", account).Group("contract_address").Fields("account,contract_address,count(contract_address) as count").All()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	var list []*apiEntity.NFTHoldInfo
	if err = result.Structs(&list); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &v1.GetNftHoldRes{
		List: list,
	}, nil
}
