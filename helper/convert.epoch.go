package helper

import (
	"time"
)

func ConvertToEpoch(date string, layout string) (int64, error) {
	t, err := time.Parse(layout, date)
	if err != nil {
		return 0, err
	}
	epoch := t.Unix()
	return epoch, nil
}
