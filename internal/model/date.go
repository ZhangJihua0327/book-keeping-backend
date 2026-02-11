package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// Date represents a date in YYYY-MM-DD format
type Date time.Time

// UnmarshalJSON parses a JSON string in YYYY-MM-DD format
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

// MarshalJSON formats the date as a JSON string in YYYY-MM-DD format
func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))), nil
}

// Value implements the driver.Valuer interface
func (d Date) Value() (driver.Value, error) {
	return time.Time(d), nil
}

// Scan implements the sql.Scanner interface
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date(time.Time{})
		return nil
	}
	if t, ok := value.(time.Time); ok {
		*d = Date(t)
		return nil
	}
	return fmt.Errorf("cannot scan type %T into Date", value)
}

// String implements the fmt.Stringer interface
func (d Date) String() string {
	return time.Time(d).Format("2006-01-02")
}
