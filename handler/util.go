package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func SetHeaders(req *http.Request, headers map[string]string) {
	req.Header.Set("Content-Type", "application/json")
	for key, val := range headers {
		req.Header.Set(key, val)
	}
}

func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func GetQueryParams(r *http.Request, key string) []string {
	return r.URL.Query()[key]
}

func ParseBody(body io.ReadCloser, dest interface{}) error {
	defer body.Close()
	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, &dest)
	if err != nil {
		return err
	}
	return nil
}

func ParseFile(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func SaveFile(body io.ReadCloser, path string, perm os.FileMode) error {
	defer body.Close()
	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, raw, perm)
	if err != nil {
		return err
	}
	return nil
}
