package global

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/v2/frame/g"
	"math/big"
	"reflect"
)

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

func GetEventLogTime(ctx context.Context, log types.Log) (uint64, error) {
	client, err := ethclient.Dial(RpcUri)
	if err != nil {
		g.Log().Error(ctx, err)
		return 0, err
	}
	header, err := client.HeaderByNumber(ctx, big.NewInt(int64(log.BlockNumber)))
	if err != nil {
		g.Log().Error(ctx, err)
		return 0, err
	}
	return header.Time, nil
}
