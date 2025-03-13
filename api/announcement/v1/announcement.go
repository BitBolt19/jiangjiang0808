package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"nez-server/internal/apiEntity"
)

type GetNewsReq struct {
	g.Meta `path:"/announcement" tags:"公告" method:"get" summary:"公告列表"`
	Type   uint `json:"type" dc:"公告类型（3 全屏公告 /  4 新闻咨询）"`
	Page   int  `json:"page" in:"query" d:"1"  v:"min:1#分页号码错误"     dc:"分页号码，默认1"`
	Size   int  `json:"size" in:"query" d:"20" v:"max:20#分页数量最大20条" dc:"分页数量，最大20"`
}

type GetNewsRes struct {
	List  []*apiEntity.AnnouncementInfo
	Total int `json:"total" dc:"总数"`
}
