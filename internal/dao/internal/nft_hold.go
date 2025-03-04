// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NftHoldDao is the data access object for the table nft_hold.
type NftHoldDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns NftHoldColumns // columns contains all the column names of Table for convenient usage.
}

// NftHoldColumns defines and stores column names for the table nft_hold.
type NftHoldColumns struct {
	Id              string //
	Account         string // 账户
	ContractAddress string // NFT合约地址
	TokenId         string // NFT tokenId
	HoldTime        string // 持有时间
	CreatedAt       string //
	UpdatedAt       string //
}

// nftHoldColumns holds the columns for the table nft_hold.
var nftHoldColumns = NftHoldColumns{
	Id:              "id",
	Account:         "account",
	ContractAddress: "contract_address",
	TokenId:         "token_id",
	HoldTime:        "hold_time",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewNftHoldDao creates and returns a new DAO object for table data access.
func NewNftHoldDao() *NftHoldDao {
	return &NftHoldDao{
		group:   "default",
		table:   "nft_hold",
		columns: nftHoldColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *NftHoldDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *NftHoldDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *NftHoldDao) Columns() NftHoldColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *NftHoldDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *NftHoldDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *NftHoldDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
