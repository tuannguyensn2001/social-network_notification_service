package tracking_transport

import (
	"context"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	trackingpb "social-work_notification_service/src/pb/tracking"
)

type transport struct {
	trackingpb.UnimplementedTrackingServiceServer
}

func New(ctx context.Context) *transport {
	t := &transport{}

	go t.GetMessageFromKafka(ctx)

	return t
}

func (t *transport) GetViewProfileByUserId(ctx context.Context, request *trackingpb.GetViewProfileByUserIdRequest) (*trackingpb.GetViewProfileByUserIdResponse, error) {
	return &trackingpb.GetViewProfileByUserIdResponse{
		Message: "done",
	}, nil
}

func (t *transport) GetMessageFromKafka(ctx context.Context) {
	go func() {
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			GroupID: "view-profile-group",
			Topic:   "view-profile",
		})

		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				zap.S().Error(err)
				break
			}
			zap.S().Info(string(m.Value))
			zap.S().Info(m)
		}
	}()
}
