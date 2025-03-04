// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RewardPoolClaimDao is the data access object for the table reward_pool_claim.
type RewardPoolClaimDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of the current DAO.
	columns RewardPoolClaimColumns // columns contains all the column names of Table for convenient usage.
}

// RewardPoolClaimColumns defines and stores column names for the table reward_pool_claim.
type RewardPoolClaimColumns struct {
	Id        string //
	Account   string // 账户
	PoolId    string // 轮次
	Token     string // 代币合约地址
	Amount    string // 数量
	TxHash    string //
	TxTime    string //
	TxEventId string //
	CreatedAt string //
	UpdatedAt string //
}

// rewardPoolClaimColumns holds the columns for the table reward_pool_claim.
var rewardPoolClaimColumns = RewardPoolClaimColumns{
	Id:        "id",
	Account:   "account",
	PoolId:    "pool_id",
	Token:     "token",
	Amount:    "amount",
	TxHash:    "tx_hash",
	TxTime:    "tx_time",
	TxEventId: "tx_event_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewRewardPoolClaimDao creates and returns a new DAO object for table data access.
func NewRewardPoolClaimDao() *RewardPoolClaimDao {
	return &RewardPoolClaimDao{
		group:   "default",
		table:   "reward_pool_claim",
		columns: rewardPoolClaimColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RewardPoolClaimDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RewardPoolClaimDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RewardPoolClaimDao) Columns() RewardPoolClaimColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RewardPoolClaimDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RewardPoolClaimDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *RewardPoolClaimDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
