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
	return &sNFT{ContractAddress: global.Buy}
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
	event, err := contract.ParseBuyNft(log.Log)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if event == nil {
		return errors.New("parse event failed")
	}
	//g.Log().Infof(ctx, "Handle nft transfer", From, event.Member.Hex())
	return s.newNFTTransfer(ctx, contractAddress, event, log)
}

func (s *sNFT) newNFTTransfer(ctx context.Context, contractAddress common.Address, transfer *buy.BuyBuyNft, log *scanner.Elog) error {
	hash := log.TxHash.Hex()
	rec, err := dao.NftTransfer.Ctx(ctx).One("tx_hash = ? AND tx_event_id = ? ", hash, log.Index)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if !rec.IsEmpty() {
		return nil
	}
	contract := contractAddress.Hex()
	from := "0x0001"
	to := transfer.Member.Hex()
	//tokenId := transfer.TokenId.Uint64()
	eventLogTime, err := global.GetEventLogTime(ctx, log.Log)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	res, err := dao.NftTransfer.Ctx(ctx).Insert(&entity.NftTransfer{
		Account:         to,
		ContractAddress: contract,
		FromAddress:     from,
		ToAddress:       to,
		//TokenId:         uint(tokenId),
		TxHash:    hash,
		TxEventId: log.Index,
		TxTime:    gtime.NewFromTimeStamp(int64(eventLogTime)),
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
	return service.NFTHold().NewTransfer(ctx, contract, to, 0, time.Unix(int64(eventLogTime), 0))
}

func (s *sNFT) GetLastTransfer(ctx context.Context, contractAddress string, to string, tokenId uint64) (*entity.NftTransfer, error) {
	lastCreatedTime, err := dao.NftTransfer.Ctx(ctx).Where("contract_address = ? AND to_address = ? AND token_id = ?", contractAddress, to, tokenId).Max("created_at")
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	rec, err := dao.NftTransfer.Ctx(ctx).One("created_at", lastCreatedTime)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	var NFTTransfer entity.NftTransfer
	err = rec.Struct(&NFTTransfer)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &NFTTransfer, nil
}
