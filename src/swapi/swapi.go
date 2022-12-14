package swapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"starwars/src/entity"
	"starwars/src/logger"
)

func GetAppearances(name string) int {
	response, err := http.Get("https://swapi.dev/api/planets?search=" + name)
	if err != nil {
		logger.Log().Warn("%v", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Log().Warn("%v", err)
	}
	var responseObject entity.ApiResponse
	json.Unmarshal(responseData, &responseObject)

	if len(responseObject.Result) > 0 {
		return len(responseObject.Result[0].Films)
	}

	return 0
}
