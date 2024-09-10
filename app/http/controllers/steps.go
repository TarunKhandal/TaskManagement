package controllers

import (
	behaviourSteps "TaskManagement/app/behaviour/steps"
	"TaskManagement/app/helpers"
	"TaskManagement/app/http/responses"
	repoSteps "TaskManagement/app/repository/steps"
	"errors"
	"net/http"
	"strconv"

	"github.com/gookit/validate"
)

func GetStepsByTaskID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("taskId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, err.Error())
		return
	}
	if id <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	objs, cnt, err := behaviourSteps.NewStepsRepository().GetSteps(id)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", nil, err.Error())
		return
	}

	if cnt == 0 {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", map[string]interface{}{
			"steps": objs, "count": cnt}, nil)
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data requested is successfully fetched", map[string]interface{}{
		"steps": objs, "count": cnt}, nil)
}

func GetStepByTaskID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("taskId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, err.Error())
		return
	}
	if id <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	sid, err := strconv.Atoi(req.PathValue("stepId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Step Id not provided", nil, err.Error())
		return
	}

	if sid <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	objs, cnt, err := behaviourSteps.NewStepsRepository().GetStep(id, sid)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", nil, err.Error())
		return
	}

	if cnt == 0 {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", map[string]interface{}{
			"steps": objs, "count": cnt}, nil)
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data requested is successfully fetched", map[string]interface{}{
		"step": objs, "count": cnt}, nil)
}

func AddStepsByTaskID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("taskId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, err.Error())
		return
	}

	if id <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	var obj repoSteps.Step
	err = helpers.BindingPayload(req, &obj)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data malformed", nil, err.Error())
		return
	}

	// overwrite the task id
	obj.TaskID = id
	v := validate.Struct(&obj)
	if !v.Validate() {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data malformed", nil, v.Errors)
		return

	}

	insertedId, err := behaviourSteps.NewStepsRepository().AddSteps(obj)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not added", nil, err.Error())
		return
	}
	if _, ok := insertedId.(int64); !ok {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not added", nil, nil)
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusCreated, "Data successfully added ", map[string]interface{}{
		"stepId": insertedId}, nil)

}

func UpdateStepsByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("taskId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Task Id not provided", nil, err.Error())
		return
	}

	if id <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	sid, err := strconv.Atoi(req.PathValue("stepId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Step Id not provided", nil, err.Error())
		return
	}

	if sid <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	var reqPlayload = make(map[string]interface{})
	err = helpers.BindingPayload(req, &reqPlayload)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data malformed", nil, err.Error())
		return
	}

	if reqPlayload["step"] != nil && len(reqPlayload["step"].(string)) == 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data malformed", nil, errors.New("step name cannot be empty").Error())
		return
	}

	// overwriting the id if it comes in the payload
	reqPlayload["task_id"] = id
	reqPlayload["id"] = sid
	err = behaviourSteps.NewStepsRepository().UpdateSteps(reqPlayload)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not updated", nil, err.Error())
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data successfully added ", nil, nil)

}

func DeleteStepsByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("taskId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Task Id not provided", nil, err.Error())
		return
	}
	if id <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	sid, err := strconv.Atoi(req.PathValue("stepId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Step Id not provided", nil, err.Error())
		return
	}
	if sid <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	err = behaviourSteps.NewStepsRepository().DeleteSteps(sid, id)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not deleted", nil, err.Error())
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data successfully deleted ", nil, nil)
}
