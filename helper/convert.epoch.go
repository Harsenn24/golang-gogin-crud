package helper

import (
	"time"
)

func ConvertToEpoch(date string) (int64, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, err
	}
	epoch := t.Unix()
	return epoch, nil
}

func ConvertToTime(epoch int64) (string, error) {
	t := time.Unix(int64(epoch), 0)

	date_string := t.Format("2006-01-02")

	return date_string, nil
}
