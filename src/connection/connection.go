package connection

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"starwars/src/entity"
	"starwars/src/logger"
)

func ConnectApi(name string) entity.Planet {
	response, err := http.Get("https://swapi.dev/api/planets?search=" + name)
	if err != nil {
		logger.Log().Warn("%v", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Log().Warn("%v", err)
	}
	var responseObject entity.Planet
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}
