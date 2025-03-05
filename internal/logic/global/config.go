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
	GP                            common.Address
	RpcUri                        string
	NFTs                          []common.Address
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
	if gVer, err := g.Cfg().Get(ctx, "contract.address.invite"); err != nil {
		panic(err)
	} else {
		InviteContract = common.HexToAddress(gVer.String())
	}
	if gVer, err := g.Cfg().Get(ctx, "contract.address.GP"); err != nil {
		panic(err)
	} else {
		GP = common.HexToAddress(gVer.String())
	}
	if gVer, err := g.Cfg().Get(ctx, "contract.address.NFTs"); err != nil {
		panic(err)
	} else {
		if err = gVer.Structs(&NFTs); err != nil {
			panic(err)
		}
	}
	if gVer, err := g.Cfg().Get(ctx, "contract.address.NFTStaking"); err != nil {
		panic(err)
	} else {
		NFTStakingContract = common.HexToAddress(gVer.String())
	}
	if gVer, err := g.Cfg().Get(ctx, "contract.address.USDT"); err != nil {
		panic(err)
	} else {
		USDT = common.HexToAddress(gVer.String())
	}
	if gVer, err := g.Cfg().Get(ctx, "contract.address.POWER"); err != nil {
		panic(err)
	} else {
		POWER = gVer.String()
	}
	if gVer, err := g.Cfg().Get(ctx, "contract.address.NFTTokenWithdraw"); err != nil {
		panic(err)
	} else {
		NFTTokenWithdraw = common.HexToAddress(gVer.String())
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
