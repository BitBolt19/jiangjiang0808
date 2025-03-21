package announcement

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "nez-server/api/announcement/v1"
	"nez-server/internal/apiEntity"
	"nez-server/internal/dao"
	"nez-server/internal/service"
)

type (
	sAnnouncement struct{}
)

func init() {
	service.RegisterAnnouncement(NewAnnouncement())
}

func NewAnnouncement() service.IAnnouncement {
	return &sAnnouncement{}
}

// 获取 对应的公告内容  列表
func (s *sAnnouncement) GetAnnouncement(ctx context.Context, req *v1.GetNewsReq) (res *v1.GetNewsRes, total int, err error) {
	typeId := req.Type
	language := req.Language
	if language == "" {
		language = "cn"
	}
	result, total, err := dao.Announcement.Ctx(ctx).Fields("id,type,language,title,link,article,index_img,start_time,end_time,status,created_at,updated_at").Where("type = ?", typeId).Where("language = ?", language).Where("status = 0").Page(req.Page, req.Size).AllAndCount(false)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, total, err
	}
	var list []*apiEntity.AnnouncementInfo
	if err = result.Structs(&list); err != nil {
		g.Log().Error(ctx, err)
		return nil, total, err
	}
	return &v1.GetNewsRes{
		List: list,
	}, total, nil
}

// 获取 公告详情
func (s *sAnnouncement) GetAnnouncementArt(ctx context.Context, req *v1.GetNewsArtReq) (res *apiEntity.AnnouncementArtInfo, err error) {
	Id := req.ID
	res = &apiEntity.AnnouncementArtInfo{}
	err = dao.Announcement.Ctx(ctx).Fields("id,type,language,title,link,article,index_img,start_time,end_time,status,created_at,updated_at").Where("id = ?", Id).Where("status = 0").Scan(res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return res, nil
}
