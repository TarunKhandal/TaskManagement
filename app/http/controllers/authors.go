package controllers

import (
	behaviourAuthor "TaskManagement/app/behaviour/author"
	"TaskManagement/app/helpers"
	"TaskManagement/app/http/responses"
	repoAuthor "TaskManagement/app/repository/author"
	"errors"
	"net/http"
	"strconv"

	"github.com/gookit/validate"
)

func GetAllAuthor(res http.ResponseWriter, req *http.Request) {

	objs, cnt, err := behaviourAuthor.NewAuthorRepository().GetAllAuthor()
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", nil, err.Error())
		return
	}

	if cnt == 0 {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", map[string]interface{}{
			"authors": objs, "count": cnt}, nil)
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data requested is successfully fetched", map[string]interface{}{
		"authors": objs, "count": cnt}, nil)

}

func GetAuthorByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("authorId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, err.Error())
		return
	}
	if id <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	obj, err := behaviourAuthor.NewAuthorRepository().GetAuthor(id)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", nil, err.Error())
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data requested is successfully fetched", obj, nil)

}

func AddAuthor(res http.ResponseWriter, req *http.Request) {

	var obj repoAuthor.Author
	err := helpers.BindingPayload(req, &obj)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data malformed", nil, err.Error())
		return
	}

	v := validate.Struct(&obj)
	if !v.Validate() {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data malformed", nil, v.Errors)
		return

	}

	insertedId, err := behaviourAuthor.NewAuthorRepository().AddAuthor(obj)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not added", nil, err.Error())
		return
	}
	if _, ok := insertedId.(int64); !ok {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not added", nil, nil)
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusCreated, "Data successfully added ", map[string]interface{}{
		"authorId": insertedId}, nil)

}

func UpdateAuthorByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("authorId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, err.Error())
		return
	}

	if id <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	var reqPlayload struct {
		FirstName interface{} `json:"first_name" filter:"Namify" validate:"string|maxLen:25"`
		LastName  interface{} `json:"last_name" filter:"Namify" validate:"string|maxLen:25"`
	}

	err = helpers.BindingPayload(req, &reqPlayload)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data malformed", nil, err.Error())
		return
	}

	v := validate.Struct(&reqPlayload)
	if !v.Validate() {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data malformed", nil, v.Errors)
		return

	}

	var reqData = map[string]interface{}{}
	err = helpers.ConvertUsingJson(reqPlayload, &reqData)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data malformed", nil, err.Error())
		return
	}

	// overwriting the id if it comes in the payload
	reqData["id"] = id
	err = behaviourAuthor.NewAuthorRepository().UpdateAuthor(reqData)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not updated", nil, err.Error())
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data successfully added ", nil, nil)

}

func DeleteAuthorByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("authorId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, err.Error())
		return
	}

	if id <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	err = behaviourAuthor.NewAuthorRepository().DeleteAuthor(id)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not deleted", nil, err.Error())
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data successfully deleted ", nil, nil)
}
