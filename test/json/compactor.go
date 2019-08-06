package json

import (
	"bytes"
	"encoding/json"
)

func Compact(uncompactedJson string) string {
	jsonBytes := []byte(uncompactedJson)
	buffer := new(bytes.Buffer)

	if err := json.Compact(buffer, jsonBytes); err != nil {
		panic(err)
	}

	return string(buffer.Bytes())
}
