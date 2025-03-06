package nft

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethscanner/ethereum-log-scanner/core/scanner"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"nez-server/internal/dao"
	"nez-server/internal/logic/consumer"
	"nez-server/internal/logic/contract/buy"
	"nez-server/internal/logic/global"
	"nez-server/internal/model/entity"
	"nez-server/internal/service"
)

type sNFTBuy struct {
	ContractAddress common.Address
}

func init() {
	service.RegisterNFTBuy(NewNFTBuy())
}

func NewNFTBuy() service.INFTBuy {
	return &sNFTBuy{ContractAddress: global.Buy}
}

// ConsumeEvents 消费链上事件
func (s *sNFTBuy) ConsumeEvents(ctx context.Context) error {
	if logs, err := consumer.ConsumeAllEvent(ctx, s.ContractAddress.String()); err != nil {
		return err
	} else {
		ids := make([]uint64, 0)
		for _, log := range logs {
			err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				return s.handleEventLog(ctx, s.ContractAddress, &log)
			})
			if err != nil {
				g.Log().Errorf(ctx, "Handle event log 失败, 错误:%+v", err)
			} else {
				ids = append(ids, *log.Id)
			}
		}
		if len(ids) > 0 {
			return consumer.ACK(ctx, ids)
		}
	}
	return nil
}

func (s *sNFTBuy) handleEventLog(ctx context.Context, contractAddress common.Address, log *scanner.Elog) error {
	client, err := ethclient.Dial(global.RpcUri)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	defer client.Close()
	contract, err := buy.NewBuy(contractAddress, client)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	event, _ := contract.ParseBuyNft(log.Log)
	//if err != nil {
	//	g.Log().Error(ctx, err)
	//	return err
	//}
	if event != nil {
		return s.newBuy(ctx, event, log)
		//return errors.New("parse event failed")
	}
	//g.Log().Infof(ctx, "Handle nft buy", event.Member.Hex(), event.Token.Hex())
	return nil
}

func (s *sNFTBuy) newBuy(ctx context.Context, event *buy.BuyBuyNft, log *scanner.Elog) error {
	hash := log.TxHash.Hex()
	eventId := log.Index
	rec, err := dao.NftBuy.Ctx(ctx).One("tx_hash = ? AND tx_event_id = ?", hash, eventId)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if rec.IsEmpty() {
		time, err := global.GetEventLogTime(ctx, log.Log)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		newBuy := &entity.NftBuy{
			Account:         event.Member.Hex(),
			ContractAddress: event.Token.Hex(),
			//TokenId:         uint(event.TokenId.Uint64()),
			TxHash:    hash,
			TxTime:    gtime.NewFromTimeStamp(int64(time)),
			TxEventId: eventId,
		}
		insertId, err := dao.NftBuy.Ctx(ctx).InsertAndGetId(newBuy)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		newBuy.Id = uint(insertId)
		//第一版先不自动结算
		//return service.NFTBuyReward().HandleNewBuy(ctx, newBuy)
		//service.NFTHold().NewTransfer(ctx, event.Token.Hex(), event.Member.Hex(), 0, times.Unix(int64(time), 0))
	}
	return nil
}

func (s *sNFTBuy) GetBuyInfo(ctx context.Context, account string) (map[string]uint64, error) {
	info := make(map[string]uint64)
	res, err := dao.NftBuy.Ctx(ctx).All("account = ?", account)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if res.IsEmpty() {
		return info, nil
	}
	buyInfoList := make([]entity.NftBuy, 0)
	err = res.Structs(&buyInfoList)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	for _, nftBuy := range buyInfoList {
		info[nftBuy.ContractAddress]++
	}
	return info, nil
}
