package dto

import "github.com/adnvilla/kafka-events-go/events-api/app/application/usecase"

type TrackEventResponseDto struct {
	Id string `json:"id"`
}

func ToTrackEventResponseDto(response usecase.TrackEventOutput) TrackEventResponseDto {
	return TrackEventResponseDto{
		Id: response.Id,
	}
}
