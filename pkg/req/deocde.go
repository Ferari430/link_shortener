package req

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T // это  то что приходит от запроса на api

	err := json.NewDecoder(body).Decode(&payload) // это  то что приходит от запроса на api
	if err != nil {
		return payload, err
	}
	return payload, nil
}
