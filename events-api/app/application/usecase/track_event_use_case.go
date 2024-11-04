package usecase

import (
	"context"

	"github.com/adnvilla/kafka-events-go/events-api/app/domain/entity"
	"github.com/adnvilla/kafka-events-go/events-api/app/domain/repository"
	"github.com/adnvilla/kafka-events-go/events-api/app/domain/vo"
)

type TrackEventUseCase interface {
	Handle(ctx context.Context, input TrackEventInput) (TrackEventOutput, error)
}

type TrackEventInput struct {
	EventName string
	Data      string
}

type TrackEventOutput struct {
	Id string
}

type trackEventUseCase struct {
	r repository.TrackEventRepository
}

func NewTrackEventUseCase(r repository.TrackEventRepository) TrackEventUseCase {
	return &trackEventUseCase{r}
}

func (uc *trackEventUseCase) Handle(ctx context.Context, input TrackEventInput) (TrackEventOutput, error) {

	result, err := uc.r.TrackEvent(ctx, vo.TrackEventVo{
		Name: input.EventName,
		Data: input.Data,
	})

	if err != nil {
		return TrackEventOutput{}, err
	}

	out := ToTrackEventOutput(result)

	return out, nil
}

func ToTrackEventOutput(result entity.TrackEventEntity) TrackEventOutput {
	return TrackEventOutput{
		Id: result.Id,
	}
}
