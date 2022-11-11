package types

import "time"

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(time.RFC3339, string(b))
	if err != nil {
		return
	}
	t.Time = date
	return
}

func (t *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(`"` + "2006-01-02" + `"`)), nil
}
