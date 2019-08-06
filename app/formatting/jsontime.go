package formatting

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	time := time.Time(t)

	if time.IsZero() {
		return []byte("null"), nil
	}

	stamp := fmt.Sprintf(`"%s"`, time.Format("2006-01-02T15:04:05-0700"))

	return []byte(stamp), nil
}
