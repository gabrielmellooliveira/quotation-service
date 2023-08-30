package utils

import "encoding/json"

func ConvertToJson[T interface{}](content T) ([]byte, error) {
	return json.Marshal(content)
}

func ConvertFromJson[T interface{}](content []byte) (*T, error) {
	var data T
	err := json.Unmarshal(content, &data)

    return &data, err
}