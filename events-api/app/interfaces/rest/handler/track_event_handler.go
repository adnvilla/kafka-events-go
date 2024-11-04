package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/adnvilla/dispatcher-go"
	"github.com/adnvilla/kafka-events-go/events-api/app/application/usecase"
	"github.com/adnvilla/kafka-events-go/events-api/app/interfaces/rest/dto"
	"github.com/gin-gonic/gin"
)

type TrackEventHandler struct {
}

func NewTrackEventHandler() TrackEventHandler {
	return TrackEventHandler{}
}

func (handler *TrackEventHandler) TrackEvent(c *gin.Context) {
	var body dto.TrackEventRequestDto
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(errors.New("please check the request")))
		return
	}

	input := usecase.TrackEventInput{
		EventName: body.Topic,
		Data:      body.Data,
	}

	result, err := dispatcher.Send[usecase.TrackEventInput, usecase.TrackEventOutput](c, input)

	if err != nil {
		// Errs it will be customize with handle errors
		c.JSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}

	c.JSON(http.StatusOK, dto.ToTrackEventResponseDto(result))
}
