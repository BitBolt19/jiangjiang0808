package announcement

import (
	"context"
	"nez-server/internal/service"

	"nez-server/api/announcement/v1"
)

func (c *ControllerV1) GetNews(ctx context.Context, req *v1.GetNewsReq) (res *v1.GetNewsRes, err error) {
	res, _, err = service.Announcement().GetAnnouncement(ctx, req)
	return
}
