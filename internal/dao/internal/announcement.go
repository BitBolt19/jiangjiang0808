// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AnnouncementDao is the data access object for the table announcement.
type AnnouncementDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns AnnouncementColumns // columns contains all the column names of Table for convenient usage.
}

// AnnouncementColumns defines and stores column names for the table announcement.
type AnnouncementColumns struct {
	Id        string //
	Type      string // announcement类型: 1:公告 2:帮助中心 3.首页弹出 4.banner
	Language  string // 语言类型
	Title     string // 标题
	Link      string // 链接地址
	Article   string // 文章
	IndexImg  string // 首页图
	StartTime string // 开始时间 毫秒
	EndTime   string // 结束 毫秒
	Status    string // 状态
	CreatedAt string // 创建时间 毫秒
	UpdatedAt string // 更新时间 毫秒
}

// announcementColumns holds the columns for the table announcement.
var announcementColumns = AnnouncementColumns{
	Id:        "id",
	Type:      "type",
	Language:  "language",
	Title:     "title",
	Link:      "link",
	Article:   "article",
	IndexImg:  "index_img",
	StartTime: "start_time",
	EndTime:   "end_time",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAnnouncementDao creates and returns a new DAO object for table data access.
func NewAnnouncementDao() *AnnouncementDao {
	return &AnnouncementDao{
		group:   "default",
		table:   "announcement",
		columns: announcementColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AnnouncementDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AnnouncementDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AnnouncementDao) Columns() AnnouncementColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AnnouncementDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AnnouncementDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AnnouncementDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
