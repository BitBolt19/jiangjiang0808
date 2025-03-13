// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package announcement

import (
	"context"

	"nez-server/api/announcement/v1"
)

type IAnnouncementV1 interface {
	GetNews(ctx context.Context, req *v1.GetNewsReq) (res *v1.GetNewsRes, err error)
}
