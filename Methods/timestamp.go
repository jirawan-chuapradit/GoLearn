package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) String() string {
	if ts.IsZero() {
		return "unknown"
	}

	const layout = "2006/01"
	return ts.Format(layout)
}

func toTimestamp(v interface{}) (ts timestamp) {
	var t int
	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	//Mon Jan 2 15:04:05 -0700 MST 2006

	ts.Time = time.Unix((int64(t)), 0)
	return ts
}

func (ts timestamp) MarshalJSON() (data []byte, _ error) {
	//ts >> integer      -->  ts.Unix() >> integer
	//	data << integer  -->  strconv.AppendInt(data, integer, 10)
	return strconv.AppendInt(data, ts.Unix(), 10), nil
}

func (ts *timestamp) UnmarshalJSON(data []byte) error {
	//118281600
	*ts = toTimestamp(string(data))
	return nil
}