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
	"nez-server/internal/logic/contract/staking"
	"nez-server/internal/logic/global"
	"nez-server/internal/model/entity"
	"nez-server/internal/service"
)

type sNFTStaking struct {
	ContractAddress common.Address
}

func init() {
	service.RegisterNFTStaking(NewNFTStaking())
}

func NewNFTStaking() service.INFTStaking {
	return &sNFTStaking{ContractAddress: global.Staking}
}

// ConsumeEvents 消费链上事件
func (s *sNFTStaking) ConsumeEvents(ctx context.Context) error {
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

func (s *sNFTStaking) handleEventLog(ctx context.Context, contractAddress common.Address, log *scanner.Elog) error {
	client, err := ethclient.Dial(global.RpcUri)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	defer client.Close()
	contract, err := staking.NewStaking(contractAddress, client)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	event, err := contract.ParseStaked(log.Log)
	//if err != nil {
	//	g.Log().Error(ctx, err)
	//	return err
	//}
	//if event == nil {
	//	return errors.New("parse event failed")
	//}
	//g.Log().Infof(ctx, "Handle nft staking", event.User.Hex(), event.Raw.Address.Hex())
	//return s.newStaking(ctx, s.ContractAddress, event, log)
	if event != nil {
		return s.newStaking(ctx, contractAddress, event, log)
	}
	return nil
}

func (s *sNFTStaking) newStaking(ctx context.Context, contractAddress common.Address, event *staking.StakingStaked, log *scanner.Elog) error {
	hash := log.TxHash.Hex()
	eventId := log.Index
	rec, err := dao.NftStaking.Ctx(ctx).One("tx_hash = ? AND tx_event_id = ?", hash, eventId)
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
		newBuy := &entity.NftStaking{
			Account:         event.User.Hex(),
			ContractAddress: contractAddress.String(),
			//TokenId:         uint(event.TokenId.Uint64()),
			TxHash:    hash,
			TxTime:    gtime.NewFromTimeStamp(int64(time)),
			TxEventId: eventId,
		}
		insertId, err := dao.NftStaking.Ctx(ctx).InsertAndGetId(newBuy)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		newBuy.Id = uint(insertId)
	}
	return nil
}

func (s *sNFTStaking) GetStakingInfo(ctx context.Context, account string) (map[string]uint64, error) {
	info := make(map[string]uint64)
	res, err := dao.NftStaking.Ctx(ctx).All("account = ?", account)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if res.IsEmpty() {
		return info, nil
	}
	stakingInfoList := make([]entity.NftStaking, 0)
	err = res.Structs(&stakingInfoList)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	for _, nftStaking := range stakingInfoList {
		info[nftStaking.ContractAddress]++
	}
	return info, nil
}
