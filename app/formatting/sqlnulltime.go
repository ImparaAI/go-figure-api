package formatting

import (
	"database/sql/driver"
	"time"
)

type SQLNullTime struct {
	Time  time.Time
	Valid bool
}

func (nt *SQLNullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)

	return nil
}

func (nt SQLNullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}

	return nt.Time, nil
}
