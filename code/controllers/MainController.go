package controllers

import (
	"encoding/json"
	io "io/ioutil"
	"net/http"
	"spaceclan1/spaceclan-data-gatherer/utils/errors"
)

type main_controller struct {
}

func (c main_controller) fetchUrl(url string, jsonRes interface{}) *errors.RestError {
	res, err := http.Get(url)
	if err != nil {
		restErr := errors.NewInternalRequestError(err.Error())
		return restErr
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		restErr := errors.NewInternalRequestError(err.Error())
		return restErr
	}
	err = json.Unmarshal(body, jsonRes)
	if err != nil {
		restErr := errors.NewInternalRequestError(err.Error())
		return restErr
	}
	return nil
}
