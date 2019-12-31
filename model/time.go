package model

import (
	"fmt"
	"time"
)

type Timestamp time.Time

// UnmarshalParam to use in echo query and form data
func (t *Timestamp) UnmarshalParam(src string) error {
	layout := time.RFC3339
	ts, err := time.Parse(layout, src)
	if err == nil {

		*t = Timestamp(ts)
		return nil
	}
	layout = "2006-01-02"
	ts, err = time.Parse(layout, src)
	*t = Timestamp(ts)
	return err
}

// UnmarshalJSON implement Unmarshaler for time
func (t *Timestamp) UnmarshalJSON(src []byte) error {
	layout := fmt.Sprintf(`"%s"`, time.RFC3339)
	ts, err := time.Parse(layout, string(src))
	if err == nil {

		*t = Timestamp(ts)
		return nil
	}
	layout = `"2006-01-02"`
	ts, err = time.Parse(layout, string(src))
	*t = Timestamp(ts)
	return err
}
