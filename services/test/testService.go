package test

import (
	apiErr "backend_proj_os/errors"
	testModel "backend_proj_os/models/test"
)

var data = []testModel.TestModelResponse{
	testModel.TestModelResponse{Code: "20000", Message: "Everything Alright"},
	testModel.TestModelResponse{Code: apiErr.InternalServerError.Code(), Message: apiErr.InternalServerError.EnMessage()},
}

func GetAllData() ([]testModel.TestModelResponse, error) {
	return data, nil
}
