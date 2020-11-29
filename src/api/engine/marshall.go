package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Decode(request *http.Request, output interface{}) error {
	if err := json.NewDecoder(request.Body).Decode(&output); err != nil {
		return fmt.Errorf("Something is wrong with input (%w) ", err)
	}
	return nil
}