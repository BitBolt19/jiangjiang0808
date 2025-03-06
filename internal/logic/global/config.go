package global

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Env                           string
	ScannerName                   string
	InviteContract                common.Address
	USDT                          common.Address
	POWER                         string
	RpcUri                        string
	Buy                           common.Address
	Box                           common.Address
	Lottery                       common.Address
	Staking                       common.Address
	NFTTransfer                   common.Address
	NFTStakingContract            common.Address
	NFTBuyRewards                 []*NFTBuyReward
	NFTHolderRewards              []*NFTHolderReward
	NFTHolderRewardDuration       uint64
	NFTTokenWithdraw              common.Address
	NFTTokenWithdrawSigPrivateKey *ecdsa.PrivateKey
)

type NFTBuyReward struct {
	ContractAddress common.Address //NFT合约地址
	USDTRewards     []uint         //USDT奖励
	PowerRewards    []uint         //算力奖励
}

type NFTHolderReward struct {
	ContractAddress common.Address //NFT合约地址
	Reward          uint           //GP奖励数量
}

func init() {
	ctx := gctx.GetInitCtx()
	if gVer, err := g.Cfg().Get(ctx, "env"); err != nil {
		panic(err)
	} else {
		Env = gVer.String()
	}
	if gVer, err := g.Cfg().Get(ctx, "NFTHolderRewardDuration"); err != nil {
		panic(err)
	} else {
		NFTHolderRewardDuration = gVer.Uint64()
	}
	if gVer, err := g.Cfg().Get(ctx, "scannerName"); err != nil {
		panic(err)
	} else {
		ScannerName = gVer.String()
	}
	if gVer, err := g.Cfg().Get(ctx, "contract.rpc"); err != nil {
		panic(err)
	} else {
		RpcUri = gVer.String()
	}
	if gVer, err := g.Cfg().Get(ctx, "contract.address.Buy"); err != nil {
		panic(err)
	} else {
		Buy = common.HexToAddress(gVer.String())
	}
	if gVer, err := g.Cfg().Get(ctx, "contract.address.Box"); err != nil {
		panic(err)
	} else {
		Box = common.HexToAddress(gVer.String())
	}
	if gVer, err := g.Cfg().Get(ctx, "contract.address.Lottery"); err != nil {
		panic(err)
	} else {
		Lottery = common.HexToAddress(gVer.String())
	}
	if gVer, err := g.Cfg().Get(ctx, "contract.address.Staking"); err != nil {
		panic(err)
	} else {
		Staking = common.HexToAddress(gVer.String())
	}

	if gVer, err := g.Cfg().Get(ctx, "NFTBuyRewards"); err != nil {
		panic(err)
	} else {
		if err = gVer.Structs(&NFTBuyRewards); err != nil {
			panic(err)
		}
		g.Log().Info(ctx, NFTBuyRewards)
	}
	if gVer, err := g.Cfg().Get(ctx, "NFTHolderRewards"); err != nil {
		panic(err)
	} else {
		if err = gVer.Structs(&NFTHolderRewards); err != nil {
			panic(err)
		}
		g.Log().Info(ctx, NFTBuyRewards)
	}
	//if NFTTokenWithdrawSigPrivateKeyStr, err := g.Cfg().Get(ctx, "contract.NFTTokenWithdrawSigPrivateKey"); err != nil {
	//	panic(err)
	//} else {
	//	if err != nil {
	//		g.Log().Errorf(ctx, "load NFTTokenWithdrawSigPrivateKey config err %s", err)
	//		panic(err)
	//	}
	//	NFTTokenWithdrawSigPrivateKey, err = crypto.HexToECDSA(NFTTokenWithdrawSigPrivateKeyStr.String())
	//	if err != nil {
	//		g.Log().Errorf(ctx, "set autoRewardPrivateKeyStr %s", err)
	//	}
	//}
}
