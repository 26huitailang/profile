package model_test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"profile/model"
	"reflect"
	"testing"
	"time"
)

type DS struct {
	T model.Timestamp `json:"time"`
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name string
		time string
		want model.Timestamp
	}{
		{
			name: "time format 2019-12-31",
			time: "2019-12-31",
			want: model.Timestamp(time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC)),
		},
		{
			name: "time format 2019-12-04T16:00:00.000Z",
			time: "2019-12-04T16:00:00.000Z",
			want: model.Timestamp(time.Date(2019, 12, 4, 16, 0, 0, 0, time.UTC)),
		},
		{
			name: "time format 1578425330571", // 2020/1/8 3:28:50
			time: "1578425330571",
			want: model.Timestamp(time.Date(2020, 1, 7, 19, 28, 50, 571*1e6, time.UTC)),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			data := fmt.Sprintf(`{"time":"%s"}`, tt.time)
			jsonData := []byte(data)
			var ds DS

			err := json.Unmarshal(jsonData, &ds)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.want, ds.T)
		})
	}
}

func TestTimestamp_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		t       model.Timestamp
		wantDst []byte
		wantErr bool
	}{
		{name: "ok", t: model.Timestamp(time.Date(2020, 1, 7, 19, 28, 50, 571*1e6, time.UTC)), wantDst: []byte(`1578425330571`), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDst, err := tt.t.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDst, tt.wantDst) {
				t.Errorf("MarshalJSON() gotDst = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}
