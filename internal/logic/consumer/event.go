package consumer

import (
	"context"
	"nez-server/internal/logic/global"

	"github.com/ethscanner/ethereum-log-scanner/core/scanner"
	"github.com/ethscanner/ethereum-log-scanner/core/storage"
)

var (
	EthereumLogScannerCheckStateList = []int{0, 1}
	EthereumLogScannerState          = 0
	EthereumLogScannerLimit          = 100
	logStorage                       scanner.DbLogStorage
)

func init() {
	logStorage = storage.NewGormLogStorage()
}

func ConsumeAllEvent(ctx context.Context, contractAddress string) (logs []scanner.Elog, err error) {
	query := scanner.LogQuery{
		ContractAddress: contractAddress,
		State:           &EthereumLogScannerState,
		Limit:           EthereumLogScannerLimit,
		CheckStateList:  EthereumLogScannerCheckStateList,
	}
	return logStorage.QueryLogs(ctx, query)
}

func ACK(ctx context.Context, ids []uint64) (err error) {
	return logStorage.MarkAsProcessed(ctx, global.ScannerName, ids)
}
