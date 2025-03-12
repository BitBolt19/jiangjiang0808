package lottery

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
	"nez-server/internal/logic/contract/lottery"
	"nez-server/internal/logic/global"
	"nez-server/internal/model/entity"
	"nez-server/internal/service"
)

type sRpClaim struct {
	ContractAddress common.Address
}

func init() {
	service.RegisterRpClaim(NewRpClaim())
}

func NewRpClaim() service.IRpClaim {
	return &sRpClaim{ContractAddress: global.Lottery}
}

// ConsumeEvents 消费链上事件
func (s *sRpClaim) ConsumeEvents(ctx context.Context) error {
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

func (s *sRpClaim) handleEventLog(ctx context.Context, contractAddress common.Address, log *scanner.Elog) error {
	client, err := ethclient.Dial(global.RpcUri)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	defer client.Close()
	contract, err := lottery.NewLottery(contractAddress, client)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	event, _ := contract.ParseUserWinner(log.Log)
	if event != nil {
		return s.newRpClaimlement(ctx, event, contractAddress, log)
	}
	return nil
}

func (s *sRpClaim) newRpClaimlement(ctx context.Context, event *lottery.LotteryUserWinner, contractAddress common.Address, log *scanner.Elog) error {
	hash := log.TxHash.Hex()
	eventId := log.Index
	rec, err := dao.RewardPoolClaim.Ctx(ctx).One("tx_hash = ? AND tx_event_id = ?", hash, eventId)
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
		newBox := &entity.RewardPoolClaim{
			Account:   event.User.Hex(),
			PoolId:    event.Round.TrailingZeroBits(),
			Token:     contractAddress.String(),
			TxHash:    hash,
			TxTime:    gtime.NewFromTimeStamp(int64(time)),
			TxEventId: eventId,
		}
		insertId, err := dao.RewardPoolClaim.Ctx(ctx).InsertAndGetId(newBox)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		newBox.Id = uint(insertId)
	}
	return nil
}
