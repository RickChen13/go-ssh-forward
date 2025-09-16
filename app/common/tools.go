package common

import "encoding/json"

func JsonDecode(data any) string {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonByte)
}

func JsonEncode[T any](data string) (T, error) {
	var res T
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
