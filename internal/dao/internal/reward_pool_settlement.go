// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RewardPoolSettlementDao is the data access object for the table reward_pool_settlement.
type RewardPoolSettlementDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of the current DAO.
	columns RewardPoolSettlementColumns // columns contains all the column names of Table for convenient usage.
}

// RewardPoolSettlementColumns defines and stores column names for the table reward_pool_settlement.
type RewardPoolSettlementColumns struct {
	Id        string //
	PoolId    string // 轮次
	Owners    string // 中奖人列表
	TxHash    string //
	TxTime    string //
	TxEventId string //
	CreatedAt string //
	UpdatedAt string //
}

// rewardPoolSettlementColumns holds the columns for the table reward_pool_settlement.
var rewardPoolSettlementColumns = RewardPoolSettlementColumns{
	Id:        "id",
	PoolId:    "pool_id",
	Owners:    "owners",
	TxHash:    "tx_hash",
	TxTime:    "tx_time",
	TxEventId: "tx_event_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewRewardPoolSettlementDao creates and returns a new DAO object for table data access.
func NewRewardPoolSettlementDao() *RewardPoolSettlementDao {
	return &RewardPoolSettlementDao{
		group:   "default",
		table:   "reward_pool_settlement",
		columns: rewardPoolSettlementColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RewardPoolSettlementDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RewardPoolSettlementDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RewardPoolSettlementDao) Columns() RewardPoolSettlementColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RewardPoolSettlementDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RewardPoolSettlementDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *RewardPoolSettlementDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
