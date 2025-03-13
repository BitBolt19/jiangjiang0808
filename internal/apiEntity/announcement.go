package apiEntity

type AnnouncementInfo struct {
	Id        uint   `json:"id" dc:"id"`           //
	Type      uint   `json:"type" dc:"公告类型"`       // 类型: 1:公告 2:帮助中心 3.首页弹出 4.新闻咨询
	Language  uint   `json:"language" dc:"语言类型"`   // 语言类型
	Title     string `json:"title" dc:"标题"`        // 标题
	Link      string `json:"link" dc:"链接地址"`       // 链接地址
	Article   string `json:"article" dc:"公告内容"`    // 文章
	IndexImg  string `json:"index_img" dc:"图片"`    // 首页图
	StartTime string `json:"start_time" dc:"开始时间"` // 开始时间 毫秒
	EndTime   string `json:"end_time" dc:"结束时间"`   // 结束 毫秒
	Status    string `json:"status" dc:"状态"`       // 状态
	CreatedAt string `json:"created_at" dc:"创建时间"` // 创建时间 毫秒
	UpdatedAt string `json:"updated_at" dc:"更新时间"` // 更新时间 毫秒
}
