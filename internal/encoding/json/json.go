package json

import "encoding/json"

func Marshal(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func Unmarshal(data []byte, dest interface{}) error {
	return json.Unmarshal(data, dest)
}
