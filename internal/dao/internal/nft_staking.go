// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NftStakingDao is the data access object for the table nft_staking.
type NftStakingDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns NftStakingColumns // columns contains all the column names of Table for convenient usage.
}

// NftStakingColumns defines and stores column names for the table nft_staking.
type NftStakingColumns struct {
	Id              string //
	Account         string // 账户
	ContractAddress string // NFT合约地址
	TokenId         string // NFT tokenId
	TxHash          string //
	TxTime          string //
	TxEventId       string //
	CreatedAt       string //
	UpdatedAt       string //
}

// nftStakingColumns holds the columns for the table nft_staking.
var nftStakingColumns = NftStakingColumns{
	Id:              "id",
	Account:         "account",
	ContractAddress: "contract_address",
	TokenId:         "token_id",
	TxHash:          "tx_hash",
	TxTime:          "tx_time",
	TxEventId:       "tx_event_id",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewNftStakingDao creates and returns a new DAO object for table data access.
func NewNftStakingDao() *NftStakingDao {
	return &NftStakingDao{
		group:   "default",
		table:   "nft_staking",
		columns: nftStakingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *NftStakingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *NftStakingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *NftStakingDao) Columns() NftStakingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *NftStakingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *NftStakingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *NftStakingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
