package notification_transport

import (
	"context"
	notificationpb "social-work_notification_service/src/pb/notification"
	"time"
)

type transport struct {
	notificationpb.UnimplementedNotificationServiceServer
}

func New(ctx context.Context) *transport {
	result := &transport{}

	return result
}

func (t *transport) GetByUserId(ctx context.Context, request *notificationpb.GetByUserIdRequest) (*notificationpb.GetByUserIdResponse, error) {
	time.Sleep(5 * time.Second)
	return &notificationpb.GetByUserIdResponse{
		Message: "success",
	}, nil
}
