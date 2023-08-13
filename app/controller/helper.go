package controller

import (
	"net/http"
	"strconv"
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
