package nft

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"nez-server/internal/dao"
	"nez-server/internal/model/entity"
	"nez-server/internal/service"
	"time"
)

type sNFTHold struct {
}

func init() {
	service.RegisterNFTHold(NewNFTHold())
}

func NewNFTHold() service.INFTHold {
	return &sNFTHold{}
}

func (s *sNFTHold) NewTransfer(ctx context.Context, contract string, to string, tokenId uint64, holdTime time.Time) error {
	rec, err := dao.NftHold.Ctx(ctx).One("contract_address = ? AND tx_event_id = ?", contract, tokenId)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	var NFTHold entity.NftHold
	if rec.IsEmpty() {
		NFTHold = entity.NftHold{
			Account:         to,
			ContractAddress: contract,
			TokenId:         uint(tokenId),
			HoldTime:        gtime.NewFromTime(holdTime),
		}
		res, err := dao.NftHold.Ctx(ctx).Insert(&NFTHold)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		if rows != 1 {
			err = errors.New("insert error")
			g.Log().Error(ctx, err, rows)
			return err
		}
	} else {
		err = rec.Struct(&NFTHold)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		NFTHold.Account = to
		NFTHold.HoldTime = gtime.NewFromTime(holdTime)
		res, err := dao.NftHold.Ctx(ctx).Where("contract_address = ? AND token_id = ?", NFTHold.ContractAddress, NFTHold.TokenId).Data(NFTHold).Update()
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		if rows != 1 {
			err = errors.New("update error")
			g.Log().Error(ctx, err, rows)
			return err
		}
	}
	return err
	//todo 这个地方的  return  的作用？
	//return service.NFTHolderReward().HandleNewNFTHold(ctx, &NFTHold)
}

//
//func (s *sNFTHold) GetAccountHold(ctx context.Context, req *v1.GetNftHoldReq) (*v1.GetNftHoldRes, error) {
//	account := req.Account
//	result, err := dao.NftHold.Ctx(ctx).Where("account = ?", account).Fields("account", "contract_address", "token_id").OrderDesc("created_at").All()
//	if err != nil {
//		return nil, err
//	}
//	var list []*apiEntity.NFTHoldInfo
//	if err = result.Structs(&list); err != nil {
//		return nil, err
//	}
//	return &v1.GetNftHoldRes{
//		List: list,
//	}, nil
//}

func (s *sNFTHold) GetAccountNFTHoldNum(ctx context.Context, account string) (holdNum int, err error) {
	return dao.NftHold.Ctx(ctx).Where("account = ?", account).Count()
}
