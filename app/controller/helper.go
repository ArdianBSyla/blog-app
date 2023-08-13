package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func queryParamInt(request *http.Request, name string, defaultValue int) (int, error) {
	str := request.URL.Query().Get(name)
	if str == "" {
		return defaultValue, nil
	}

	parsedVal, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return parsedVal, nil
}

// URLParam returns the url parameter as a string
// from a http.Request object.
func URLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

// URLParamInt returns the url parameter as an integer
// from a http.Request object.
func URLParamInt(r *http.Request, key string) (int, error) {
	return strconv.Atoi(URLParam(r, key))
}
