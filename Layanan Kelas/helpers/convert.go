package helper

import (
	"encoding/json"
)

func Convert(data any, target any) error {
	jsonEncode, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonEncode, &target)
	if err != nil {
		return err
	}
	return nil
}

