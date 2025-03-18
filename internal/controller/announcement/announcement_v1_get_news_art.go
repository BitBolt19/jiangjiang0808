package announcement

import (
	"context"
	"nez-server/internal/apiEntity"
	"nez-server/internal/service"

	"nez-server/api/announcement/v1"
)

func (c *ControllerV1) GetNewsArt(ctx context.Context, req *v1.GetNewsArtReq) (res *apiEntity.AnnouncementArtInfo, err error) {
	return service.Announcement().GetAnnouncementArt(ctx, req)
}
