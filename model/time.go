package model

import (
	"bytes"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"strconv"
	"time"
)

type Timestamp time.Time

func Now() Timestamp {
	now := time.Now()
	tm := now.Round(time.Millisecond)
	return Timestamp(tm)
}

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
// -> timestamp
// -> RFC3339
// -> `"2006-01-02"`
// -> return err
func (t *Timestamp) UnmarshalJSON(src []byte) error {
	src = bytes.Trim(src, "\"")
	ts, err := strconv.ParseInt(string(src), 10, 64)
	if err == nil {
		tm := time.Unix(ts/1e3, ts%1e3*1e6).UTC()
		*t = Timestamp(tm)
		return nil
	}

	layout := fmt.Sprintf(`%s`, time.RFC3339)
	tm, err := time.Parse(layout, string(src))
	if err == nil {
		*t = Timestamp(tm)
		return nil
	}

	layout = `2006-01-02`
	tm, err = time.Parse(layout, string(src))
	*t = Timestamp(tm)
	return err
}

// MarshalBSONValue time in mongo implement ValueMarshaler
func (t Timestamp) MarshalBSONValue() (_type bsontype.Type, b []byte, err error) {
	_type = bson.TypeDateTime

	//b = time.Time(t).AppendFormat(b, time.RFC3339)
	b = bsoncore.AppendTime(b, time.Time(t))
	//b, err = bson.Marshal(time.Time(t))
	return _type, b, err
}

func (t *Timestamp) UnmarshalBSONValue(_type bsontype.Type, b []byte) error {
	t2, _, _ := bsoncore.ReadTime(b)
	*t = Timestamp(t2)

	return nil
}

func (t *Timestamp) String() string {
	return time.Time(*t).String()
}
