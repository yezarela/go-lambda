package jsonutil

import (
	"bytes"
	"encoding/json"
	"io"
)

// Compact returns new json string with insignificant space characters elided
func Compact(jsonString string) (string, error) {
	jsonBytes := []byte(jsonString)
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, jsonBytes); err != nil {
		return jsonString, err
	}

	readBuf, err := io.ReadAll(buffer)
	if err != nil {
		return jsonString, err
	}

	return string(readBuf), nil
}
