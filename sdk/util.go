package sdk

import "net/http"

func IsSuccess(code int) bool {
	return code >= http.StatusOK && code < 300
}
