// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "nez-server/api/announcement/v1"
)

type (
	IAnnouncement interface {
		// 获取 对应的公告内容
		GetAnnouncement(ctx context.Context, req *v1.GetNewsReq) (res *v1.GetNewsRes, total int, err error)
	}
)

var (
	localAnnouncement IAnnouncement
)

func Announcement() IAnnouncement {
	if localAnnouncement == nil {
		panic("implement not found for interface IAnnouncement, forgot register?")
	}
	return localAnnouncement
}

func RegisterAnnouncement(i IAnnouncement) {
	localAnnouncement = i
}
