package nft

import (
	"context"
	"errors"
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
	"time"
)

type sNFT struct {
	ContractAddress common.Address
}

func init() {
	service.RegisterNFT(NewNFT())
}

func NewNFT() service.INFT {
	return &sNFT{ContractAddress: global.Box}
}

// ConsumeEvents 消费链上事件
func (s *sNFT) ConsumeEvents(ctx context.Context) error {
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

func (s *sNFT) handleEventLog(ctx context.Context, contractAddress common.Address, log *scanner.Elog) error {
	client, err := ethclient.Dial(global.RpcUri)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	defer client.Close()
	//contract, err := nft.NewNft(contractAddress, client)
	contract, err := buy.NewBuy(contractAddress, client)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	//event, err := contract.ParseTransfer(log.Log)
	event, err := contract.ParseTransferBlindBox(log.Log)
	//if err != nil {
	//	g.Log().Error(ctx, err)
	//	return err
	//}
	//if event == nil {
	//	return errors.New("parse event failed")
	//}
	if event != nil {
		return s.NewTransfer(ctx, s.ContractAddress, event.From, event.To, event.Amount.Int64(), log)
	}
	//g.Log().Infof(ctx, "Handle nft transfer", From, event.Member.Hex())
	return nil
}

func (s *sNFT) NewTransfer(ctx context.Context, contract common.Address, from common.Address, to common.Address, amount int64, log *scanner.Elog) error {
	hash := log.TxHash.Hex()
	rec, err := dao.NftTransfer.Ctx(ctx).One("tx_hash = ? AND tx_event_id = ? ", hash, log.Index)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if !rec.IsEmpty() {
		return nil
	}
	//tokenId := transfer.TokenId.Uint64()
	eventLogTime, err := global.GetEventLogTime(ctx, log.Log)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	res, err := dao.NftTransfer.Ctx(ctx).Insert(&entity.NftTransfer{
		Account:         to.Hex(),
		ContractAddress: contract.Hex(),
		FromAddress:     from.Hex(),
		ToAddress:       to.Hex(),
		Amount:          amount,
		TxHash:          hash,
		TxEventId:       log.Index,
		TxTime:          gtime.NewFromTimeStamp(int64(eventLogTime)),
	})
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
		err = errors.New("insert rows err")
		g.Log().Error(ctx, err, rows)
	}
	return service.NFTHold().NewHold(ctx, contract.Hex(), to.Hex(), uint64(log.Index), time.Unix(int64(eventLogTime), 0))
}
