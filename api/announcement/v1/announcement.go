package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"nez-server/internal/apiEntity"
)

type GetNewsReq struct {
	g.Meta   `path:"/announcement" tags:"公告" method:"get" summary:"公告列表"`
	Type     int    `json:"type" dc:"公告类型: 1:公告 2:帮助中心 3.首页弹出 4.banner 5.新闻资讯"`
	Language string `json:"language" dc:"语言类型： cn.中文 en.英文"`
	Page     int    `json:"page" in:"query" d:"1"  v:"min:1#分页号码错误"     dc:"分页号码，默认1"`
	Size     int    `json:"size" in:"query" d:"20" v:"max:20#分页数量最大20条" dc:"分页数量，最大20"`
}

type GetNewsRes struct {
	List  []*apiEntity.AnnouncementInfo
	Total int `json:"total" dc:"总数"`
}

// 公告详情
type GetNewsArtReq struct {
	g.Meta `path:"/announcement/article" tags:"公告" method:"get" summary:"公告详情"`
	ID     uint `json:"id" dc:"id"`
}
