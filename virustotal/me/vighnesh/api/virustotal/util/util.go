package util

import (
	"encoding/json"
)

func To(data []byte, typ interface{}) error {
	return json.Unmarshal(data, typ)
}
