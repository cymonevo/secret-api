package json

import "github.com/tidwall/gjson"

func get(content, path string) gjson.Result {
	return gjson.Get(content, path)
}

func Exists(content, path string) bool {
	return get(content, path).Exists()
}

func GetString(content, path string) string {
	return get(content, path).String()
}

func GetInt(content, path string) int64 {
	return get(content, path).Int()
}

func GetBool(content, path string) bool {
	return get(content, path).Bool()
}

func GetRaw(content, path string) interface{} {
	return get(content, path).Value()
}
