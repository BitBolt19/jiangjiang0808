// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Announcement is the golang structure of table announcement for DAO operations like Where/Data.
type Announcement struct {
	g.Meta    `orm:"table:announcement, do:true"`
	Id        interface{} //
	Type      interface{} // announcement类型: 1:公告 2:帮助中心 3.首页弹出 4.banner
	Language  interface{} // 语言类型
	Title     interface{} // 标题
	Link      interface{} // 链接地址
	Article   interface{} // 文章
	IndexImg  interface{} // 首页图
	StartTime interface{} // 开始时间 毫秒
	EndTime   interface{} // 结束 毫秒
	Status    interface{} // 状态
	CreatedAt interface{} // 创建时间 毫秒
	UpdatedAt interface{} // 更新时间 毫秒
}
