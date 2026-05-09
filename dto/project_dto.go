package dto

import (
	"fmt"
	"time"
)

// DateOnly is a time.Time that marshals/unmarshals JSON as "YYYY-MM-DD".
type DateOnly struct {
	time.Time
}

const dateLayout = "2006-01-02"

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := string(b)
	// strip surrounding quotes
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD: %w", err)
	}
	d.Time = t
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format(dateLayout) + `"`), nil
}

type ProjectsResponse struct {
	Id               uint     `json:"id"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Status           string   `json:"status"`
	StartDate        DateOnly `json:"start_date"`
	TargetCompletion DateOnly `json:"target_completion"`
}

type CreateProjectRequest struct {
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Status           string   `json:"status"`
	StartDate        DateOnly `json:"start_date"`
	TargetCompletion DateOnly `json:"target_completion"`
}
