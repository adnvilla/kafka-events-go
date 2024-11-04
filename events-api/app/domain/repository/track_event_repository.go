package repository

import (
	"context"

	"github.com/adnvilla/kafka-events-go/events-api/app/domain/entity"
	"github.com/adnvilla/kafka-events-go/events-api/app/domain/vo"
)

type TrackEventRepository interface {
	TrackEvent(ctx context.Context, event vo.TrackEventVo) (entity.TrackEventEntity, error)
}
