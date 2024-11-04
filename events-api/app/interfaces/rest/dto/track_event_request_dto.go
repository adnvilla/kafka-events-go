package dto

type TrackEventRequestDto struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}
