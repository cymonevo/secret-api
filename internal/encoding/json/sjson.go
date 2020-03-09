package json

import "github.com/tidwall/sjson"

func Set(content, path string, data interface{}) (string, error) {
	return sjson.Set(content, path, data)
}
