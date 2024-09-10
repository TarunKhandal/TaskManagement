package responses

import (
	"encoding/json"
	"net/http"
)

func ApiResponseGivenMsg(res http.ResponseWriter, httpCode int, msg string, data, errmsg interface{}) {
	if res == nil {
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(httpCode)

	responseData := map[string]interface{}{
		"message": msg,
		"data":    data,
		"error":   errmsg,
	}

	b, err := json.Marshal(responseData)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusUnprocessableEntity)

		response := map[string]string{
			"error": "Error Occured :" + err.Error(),
		}
		b, _ := json.Marshal(response)
		res.Write(b)
		return
	}

	res.Write(b)
	return
}
