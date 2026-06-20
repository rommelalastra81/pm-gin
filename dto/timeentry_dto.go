package dto

import (
	"fmt"
	"time"
)

// DateOnly is a time.Time that marshals/unmarshals JSON as "YYYY-MM-DD".
type DateOnlyTimeEntry struct {
	time.Time
}

type TimeOnlyTimeEntry struct {
	time.Time
}

const dateLayoutTimeEntry = "2006-01-02"
const timeLayoutTimeEntry = "15:04:05"

func (d *DateOnlyTimeEntry) UnmarshalJSON(b []byte) error {
	s := string(b)
	// handle JSON null — leave time as zero value
	if s == "null" {
		return nil
	}
	// strip surrounding quotes
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	t, err := time.Parse(dateLayoutTimeEntry, s)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD: %w", err)
	}
	d.Time = t
	return nil
}

func (d *TimeOnlyTimeEntry) UnmarshalJSON(b []byte) error {
	s := string(b)
	// handle JSON null — leave time as zero value
	if s == "null" {
		return nil
	}
	// strip surrounding quotes
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	t, err := time.Parse(timeLayoutTimeEntry, s)
	if err != nil {
		return fmt.Errorf("invalid time format, expected HH:MM:SS: %w", err)
	}
	d.Time = t
	return nil
}

func (d DateOnlyTimeEntry) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format(dateLayoutTimeEntry) + `"`), nil
}

type TimeEntryResponse struct {
	Id         uint              `json:"id" gorm:"primaryKey"`
	UserId     uint              `json:"user_id"`
	TaskId     uint              `json:"task_id"`
	Activities string            `json:"activities"`
	Date       DateOnlyTimeEntry `json:"date"`
	StartTime  TimeOnlyTimeEntry `json:"start_time"`
	EndTime    TimeOnlyTimeEntry `json:"end_time"`
}

type CreateTimeEntryRequest struct {
	UserId     uint              `json:"user_id"`
	TaskId     uint              `json:"task_id"`
	Activities string            `json:"activities"`
	Date       DateOnlyTimeEntry `json:"date"`
	StartTime  TimeOnlyTimeEntry `json:"start_time"`
	EndTime    TimeOnlyTimeEntry `json:"end_time"`
}

type UpdateTimeEntryRequest struct {
	Activities string            `json:"activities"`
	Date       DateOnlyTimeEntry `json:"date"`
	StartTime  TimeOnlyTimeEntry `json:"start_time"`
	EndTime    TimeOnlyTimeEntry `json:"end_time"`
}
