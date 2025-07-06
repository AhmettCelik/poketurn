package api

import (
	"encoding/json"
	"fmt"
	"io"
)

func decodeResponse[T any](data io.ReadCloser) (T, error) {
	var params T

	decoder := json.NewDecoder(data)
	if err := decoder.Decode(&params); err != nil {
		fmt.Printf("Error decoding json: %v\n", err)
		return params, err
	}

	return params, nil
}
