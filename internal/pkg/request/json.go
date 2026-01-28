package request

import (
	"io"
	"net/http"

	"github.com/bytedance/sonic"
)

func BindJSON[T any](r *http.Request) (T, error) {
	var result T

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return result, err
	}

	err = sonic.Unmarshal(body, &result)
	return result, err
}
