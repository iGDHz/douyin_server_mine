package service

import "time"

type Video struct {
	Video_id               int       `json:"video_id"`
	Video_location         string    `json:"video_location"`
	Video_picture_location string    `json:"video_picture_location"`
	Video_latest_time      time.Time `json:"video_latest_time"`
	Video_state            int       `json:"video_state"`
	Video_category         string    `json:"video_category"`
	Video_title            string    `json:"video_title"`
}
