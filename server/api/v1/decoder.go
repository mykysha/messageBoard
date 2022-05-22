package v1

import (
	"encoding/json"
	"fmt"
)

func Decode(data []byte, object any) error {
	err := json.Unmarshal(data, object)
	if err != nil {
		return fmt.Errorf("decoding object failure: %w", err)
	}

	return nil
}
