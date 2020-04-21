package models

import (
	"strings"
	"time"
)

type SSTime struct{ Time time.Time }

//UnmarshalJSON formats JSON input
func (m *SSTime) UnmarshalJSON(b []byte) error {
	strInput := string(b)
	strInput = strings.Trim(strInput, `"`)

	newTime, err := time.Parse(TimeFormat(), strInput)
	if err != nil {
		return err
	}

	m.Time = newTime

	return nil
}

//MarshalJSON formats JSON output
func (m SSTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + m.Time.Format(TimeFormat()) + `"`), nil
}
