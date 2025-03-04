// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccountTokenBalanceLogDao is the data access object for the table account_token_balance_log.
type AccountTokenBalanceLogDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of the current DAO.
	columns AccountTokenBalanceLogColumns // columns contains all the column names of Table for convenient usage.
}

// AccountTokenBalanceLogColumns defines and stores column names for the table account_token_balance_log.
type AccountTokenBalanceLogColumns struct {
	Id         string //
	OrderId    string // 关联订单ID
	Account    string // 账户
	Token      string // 代币
	Amount     string // 数量
	OldBalance string // 旧余额
	NewBalance string // 新余额
	Type       string // 流水类型
	CreatedAt  string //
	UpdatedAt  string //
}

// accountTokenBalanceLogColumns holds the columns for the table account_token_balance_log.
var accountTokenBalanceLogColumns = AccountTokenBalanceLogColumns{
	Id:         "id",
	OrderId:    "order_id",
	Account:    "account",
	Token:      "token",
	Amount:     "amount",
	OldBalance: "old_balance",
	NewBalance: "new_balance",
	Type:       "type",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewAccountTokenBalanceLogDao creates and returns a new DAO object for table data access.
func NewAccountTokenBalanceLogDao() *AccountTokenBalanceLogDao {
	return &AccountTokenBalanceLogDao{
		group:   "default",
		table:   "account_token_balance_log",
		columns: accountTokenBalanceLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AccountTokenBalanceLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AccountTokenBalanceLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AccountTokenBalanceLogDao) Columns() AccountTokenBalanceLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AccountTokenBalanceLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AccountTokenBalanceLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AccountTokenBalanceLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
