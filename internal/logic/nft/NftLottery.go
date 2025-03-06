package nft

//
//import (
//	"context"
//	"errors"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/ethclient"
//	"github.com/ethscanner/ethereum-log-scanner/core/scanner"
//	"github.com/gogf/gf/v2/database/gdb"
//	"github.com/gogf/gf/v2/frame/g"
//	"github.com/gogf/gf/v2/os/gtime"
//	"nez-server/internal/dao"
//	"nez-server/internal/logic/consumer"
//	"nez-server/internal/logic/contract/buy"
//	"nez-server/internal/logic/global"
//	"nez-server/internal/model/entity"
//	"nez-server/internal/service"
//)
//
//type sNFTLottery struct {
//	ContractAddress common.Address
//}
//
//func init() {
//	service.RegisterNFTLottery(NewNFTLottery())
//}
//
//func NewNFTLottery() service.INFTLottery {
//	return &sNFTLottery{ContractAddress: global.Buy}
//}
//
//// ConsumeEvents 消费链上事件
//func (s *sNFTLottery) ConsumeEvents(ctx context.Context) error {
//	if logs, err := consumer.ConsumeAllEvent(ctx, s.ContractAddress.String()); err != nil {
//		return err
//	} else {
//		ids := make([]uint64, 0)
//		for _, log := range logs {
//			err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
//				return s.handleEventLog(ctx, s.ContractAddress, &log)
//			})
//			if err != nil {
//				g.Log().Errorf(ctx, "Handle event log 失败, 错误:%+v", err)
//			} else {
//				ids = append(ids, *log.Id)
//			}
//		}
//		if len(ids) > 0 {
//			return consumer.ACK(ctx, ids)
//		}
//	}
//	return nil
//}
//
//func (s *sNFTLottery) handleEventLog(ctx context.Context, contractAddress common.Address, log *scanner.Elog) error {
//	client, err := ethclient.Dial(global.RpcUri)
//	if err != nil {
//		g.Log().Error(ctx, err)
//		return err
//	}
//	defer client.Close()
//	contract, err := buy.NewBuy(contractAddress, client)
//	if err != nil {
//		g.Log().Error(ctx, err)
//		return err
//	}
//	event, err := contract.ParseBuyNft(log.Log)
//	if err != nil {
//		g.Log().Error(ctx, err)
//		return err
//	}
//	if event == nil {
//		return errors.New("parse event failed")
//	}
//	g.Log().Infof(ctx, "Handle nft buy", event.Member.Hex(), event.Token.Hex())
//	return s.newBuy(ctx, event, log)
//}
//
//func (s *sNFTLottery) newBuy(ctx context.Context, event *buy.BuyBuyNft, log *scanner.Elog) error {
//	hash := log.TxHash.Hex()
//	eventId := log.Index
//	rec, err := dao.NFTLottery.Ctx(ctx).One("tx_hash = ? AND tx_event_id = ?", hash, eventId)
//	if err != nil {
//		g.Log().Error(ctx, err)
//		return err
//	}
//	if rec.IsEmpty() {
//		time, err := global.GetEventLogTime(ctx, log.Log)
//		if err != nil {
//			g.Log().Error(ctx, err)
//			return err
//		}
//		newBuy := &entity.NFTLottery{
//			Account:         event.Member.Hex(),
//			ContractAddress: event.Token.Hex(),
//			//TokenId:         uint(event.TokenId.Uint64()),
//			TxHash:    hash,
//			TxTime:    gtime.NewFromTimeStamp(int64(time)),
//			TxEventId: eventId,
//		}
//		insertId, err := dao.NFTLottery.Ctx(ctx).InsertAndGetId(newBuy)
//		if err != nil {
//			g.Log().Error(ctx, err)
//			return err
//		}
//		newBuy.Id = uint(insertId)
//		//第一版先不自动结算
//		//return service.NFTLotteryReward().HandleNewBuy(ctx, newBuy)
//	}
//	return nil
//}
//
//func (s *sNFTLottery) GetBuyInfo(ctx context.Context, account string) (map[string]uint64, error) {
//	info := make(map[string]uint64)
//	res, err := dao.NFTLottery.Ctx(ctx).All("account = ?", account)
//	if err != nil {
//		g.Log().Error(ctx, err)
//		return nil, err
//	}
//	if res.IsEmpty() {
//		return info, nil
//	}
//	buyInfoList := make([]entity.NFTLottery, 0)
//	err = res.Structs(&buyInfoList)
//	if err != nil {
//		g.Log().Error(ctx, err)
//		return nil, err
//	}
//	for _, NFTLottery := range buyInfoList {
//		info[NFTLottery.ContractAddress]++
//	}
//	return info, nil
//}
