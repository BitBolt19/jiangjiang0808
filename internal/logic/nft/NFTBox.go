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
	"nez-server/internal/logic/consts"
	"nez-server/internal/logic/consumer"
	"nez-server/internal/logic/contract/buy"
	"nez-server/internal/logic/global"
	"nez-server/internal/model/entity"
	"nez-server/internal/service"
	"sync"
)

type sNFTbox struct {
	sync.Mutex
	ContractAddress common.Address
}

func init() {
	service.RegisterNFTbox(NewNFTbox())
}

func NewNFTbox() service.INFTbox {
	return &sNFTbox{ContractAddress: global.Box}
}

// ConsumeEvents 消费链上事件
func (s *sNFTbox) ConsumeEvents(ctx context.Context) error {
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

func (s *sNFTbox) handleEventLog(ctx context.Context, contractAddress common.Address, log *scanner.Elog) error {
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
	//event, _ := contract.ParseBlindBoxNft(log.Log)
	event, _ := contract.ParseOpenBlindBox(log.Log)
	//if err != nil {
	//	g.Log().Error(ctx, err)
	//	return err
	//}
	//if event == nil {
	//	return errors.New("parse event failed")
	//}
	//g.Log().Infof(ctx, "Handle nft Box Open", event.Member.Hex(), event.Token.Hex())
	//return s.newBox(ctx, event, log)
	if event != nil {
		return s.newBox(ctx, event, log)
	}
	return nil
}

func (s *sNFTbox) newBox(ctx context.Context, event *buy.BuyOpenBlindBox, log *scanner.Elog) error {
	hash := log.TxHash.Hex()
	eventId := log.Index
	//查询判断数据是否已经写入
	rec, err := dao.NftBox.Ctx(ctx).LockUpdate().One("tx_hash = ? AND tx_event_id = ?", hash, eventId)
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
		newBox := &entity.NftBox{
			Account:         event.Member.Hex(),
			ContractAddress: event.Token.Hex(),
			Amount:          event.Amount.Int64(),
			TxHash:          hash,
			TxTime:          gtime.NewFromTimeStamp(int64(time)),
			TxEventId:       eventId,
		}
		insertId, err := dao.NftBox.Ctx(ctx).InsertAndGetId(newBox)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		newBox.Id = uint(insertId)
		// 调用更新状态
		err = s.UpdateStatus(ctx, event, eventId)
		if err != nil {
			return err
		}
	}
	return nil
}

// 当盲盒打开后，购买 状态 要更改为已开启
func (s *sNFTbox) UpdateStatus(ctx context.Context, event *buy.BuyOpenBlindBox, eventId uint) (err error) {
	_, err = dao.NftBuy.Ctx(ctx).
		Where("account = ?", event.Member.Hex()).
		Where("tx_event_id = ?", eventId).
		Data("status", consts.NFT_Open_Box_Status).
		Update()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return err
}
