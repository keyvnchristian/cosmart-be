package model

import "time"

type Schedule struct {
	UserID     int32     `json:"user_id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Edition    int32     `json:"edition"`
	PickUpTime time.Time `json:"pick_up_time"`
}

type ScheduleRequest struct {
	UserID         int32  `json:"user_id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Edition        int32  `json:"edition"`
	PickupSchedule string `json:"pickup_schedule"`
}
