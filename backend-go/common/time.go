package common

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Date time.Time

func NowDate() Date {
	return Date(time.Now())
}

func (d *Date) String() string {
	return time.Time(*d).Format("2006-01-02")
}

func (d *Date) Scan(value interface{}) error {
	date, err := time.Parse(time.RFC3339, value.(time.Time).Format(time.RFC3339))
	if err != nil {
		return err
	}
	*d = Date(date)
	return nil
}

func (d *Date) Value() (driver.Value, error) {
	return d.String(), nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.String())), nil

}

func (d *Date) UnmarshalJSON(data []byte) error {
	//var dateStr string
	//if err := json.Unmarshal(data, &dateStr); err != nil {
	//	return err
	//}
	dateStr := strings.Replace(string(data), "\"", "", -1)
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}

	*d = Date(parsedDate)
	return nil
}
func (d *Date) Marshal() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) Unmarshal(data []byte) error {
	dateStr := strings.Replace(string(data), "\"", "", -1)
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	*d = Date(parsedDate)
	return nil
}
func (d *Date) MarshalMsgpack() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *Date) UnmarshalMsgpack(data []byte) error {
	dateStr := strings.Replace(string(data), "\"", "", -1)
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	*d = Date(parsedDate)
	return nil
}

type Time time.Time

func (t *Time) ToTime() time.Time {
	return time.Time(*t)
}

func (t *Time) Add(d time.Duration) Time {
	return Time(t.ToTime().Add(d))
}

func (t *Time) String() string {
	return time.Time(*t).Format("15:04:05")
}

func (t *Time) Scan(value interface{}) error {
	parse, err := time.Parse("15:04:05", value.(string))
	if err != nil {
		return err
	}
	*t = Time(parse)
	return nil
}

func (t *Time) Value() (driver.Value, error) {
	return t.String(), nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	// Trim the quotes from the incoming byte slice and convert it to a string
	timeStr := strings.Replace(string(data), "\"", "", -1)

	// Parse the time string into a time.Time type
	parsedTime, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		return err
	}

	// Convert the time.Time type to a Time type and assign it to t
	*t = Time(parsedTime)

	return nil
}
