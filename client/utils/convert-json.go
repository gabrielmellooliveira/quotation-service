package utils

import "encoding/json"

func ConvertFromJson[T interface{}](content []byte) (*T, error) {
	var data T
	err := json.Unmarshal(content, &data)

    return &data, err
}