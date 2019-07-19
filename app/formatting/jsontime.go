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

    stamp := fmt.Sprintf(`"%s"`, time.Format("2010-01-01 12:12:12"))

    return []byte(stamp), nil
}