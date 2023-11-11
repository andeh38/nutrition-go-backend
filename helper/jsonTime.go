package helper

import (
	"time"
)

const (
	INPUT_FORMAT  = "2006-01-02T15:04:05.999999999-07:00"
	OUTPUT_FORMAT = "2006-01-02T15:04:05.000Z"
)

type CustomTime struct {
	time.Time
}

func UnmarshalJSON(JSONdate string) (date time.Time, err error) {
	date, err = time.Parse(OUTPUT_FORMAT, JSONdate)
	date.Format(time.RFC3339)
	if err != nil {
		return time.Time{}, err
	}
	
	return date, nil
}