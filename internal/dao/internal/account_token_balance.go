// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccountTokenBalanceDao is the data access object for the table account_token_balance.
type AccountTokenBalanceDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of the current DAO.
	columns AccountTokenBalanceColumns // columns contains all the column names of Table for convenient usage.
}

// AccountTokenBalanceColumns defines and stores column names for the table account_token_balance.
type AccountTokenBalanceColumns struct {
	Id        string //
	Account   string // 账户
	Token     string // 代币
	Balance   string // 余额
	CreatedAt string //
	UpdatedAt string //
}

// accountTokenBalanceColumns holds the columns for the table account_token_balance.
var accountTokenBalanceColumns = AccountTokenBalanceColumns{
	Id:        "id",
	Account:   "account",
	Token:     "token",
	Balance:   "balance",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAccountTokenBalanceDao creates and returns a new DAO object for table data access.
func NewAccountTokenBalanceDao() *AccountTokenBalanceDao {
	return &AccountTokenBalanceDao{
		group:   "default",
		table:   "account_token_balance",
		columns: accountTokenBalanceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AccountTokenBalanceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AccountTokenBalanceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AccountTokenBalanceDao) Columns() AccountTokenBalanceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AccountTokenBalanceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AccountTokenBalanceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AccountTokenBalanceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
