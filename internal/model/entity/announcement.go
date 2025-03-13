// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Announcement is the golang structure for table announcement.
type Announcement struct {
	Id        uint   `json:"id"        orm:"id"         ` //
	Type      int    `json:"type"      orm:"type"       ` // announcement类型: 1:公告 2:帮助中心 3.首页弹出 4.banner
	Language  string `json:"language"  orm:"language"   ` // 语言类型
	Title     string `json:"title"     orm:"title"      ` // 标题
	Link      string `json:"link"      orm:"link"       ` // 链接地址
	Article   string `json:"article"   orm:"article"    ` // 文章
	IndexImg  string `json:"indexImg"  orm:"index_img"  ` // 首页图
	StartTime int64  `json:"startTime" orm:"start_time" ` // 开始时间 毫秒
	EndTime   int64  `json:"endTime"   orm:"end_time"   ` // 结束 毫秒
	Status    int    `json:"status"    orm:"status"     ` // 状态
	CreatedAt int64  `json:"createdAt" orm:"created_at" ` // 创建时间 毫秒
	UpdatedAt int64  `json:"updatedAt" orm:"updated_at" ` // 更新时间 毫秒
}
