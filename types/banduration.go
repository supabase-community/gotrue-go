package types

import (
	"encoding/json"
	"time"
)

type BanDuration struct {
	d *time.Duration
}

func BanDurationNone() BanDuration {
	return BanDuration{
		d: nil,
	}
}

func BanDurationTime(d time.Duration) BanDuration {
	return BanDuration{
		d: &d,
	}
}

func (b BanDuration) Value() *time.Duration {
	return b.d
}

func (b BanDuration) String() string {
	if b.d == nil {
		return "none"
	}
	return b.d.String()
}

func (b BanDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *BanDuration) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "none" {
		b.d = nil
		return nil
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	b.d = &d
	return nil
}
