package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func BindingPayload(req *http.Request, obj interface{}) error {
	responseData, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	if len(responseData) == 0 {
		return errors.New("no data provided")
	}

	err = json.Unmarshal(responseData, &obj)
	if err != nil {
		return err
	}

	return nil
}

func ConvertUsingJson(src interface{}, dest interface{}) error {
	dataBytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataBytes, dest)
	if err != nil {
		return err
	}
	return nil
}

func RemoveNilKeys[K comparable](m map[K]interface{}) {
	for k, v := range m {
		if v == nil {
			delete(m, k)
		}
	}
}
