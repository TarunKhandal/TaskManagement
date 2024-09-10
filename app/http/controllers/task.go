package controllers

import (
	behaviourTask "TaskManagement/app/behaviour/task"
	"TaskManagement/app/helpers"
	"TaskManagement/app/http/responses"
	repoTask "TaskManagement/app/repository/task"
	"errors"
	"net/http"
	"strconv"

	"github.com/gookit/validate"
)

func GetAllTasks(res http.ResponseWriter, req *http.Request) {

	objs, cnt, err := behaviourTask.NewTaskRepository().GetAllTask()
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", nil, err.Error())
		return
	}

	if cnt == 0 {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", map[string]interface{}{
			"tasks": objs, "count": cnt}, nil)
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data requested is successfully fetched", map[string]interface{}{
		"tasks": objs, "count": cnt}, nil)

}

func GetTasksByAuthorID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("authorId"))
	if err != nil || id == 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Id not provided", nil, err.Error())
		return
	}

	if id <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	objs, cnt, err := behaviourTask.NewTaskRepository().GetTaskByAuthorId(id)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", nil, err.Error())
		return
	}

	if cnt == 0 {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", map[string]interface{}{
			"tasks": objs, "count": cnt}, nil)
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data requested is successfully fetched", map[string]interface{}{
		"tasks": objs, "count": cnt}, nil)

}

func GetTasksByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("taskId"))
	if err != nil || id == 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Id not provided", nil, err.Error())
		return
	}

	if id <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	obj, err := behaviourTask.NewTaskRepository().GetTask(id)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusNotFound, "Data requested is not avaliable", nil, err.Error())
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data requested is successfully fetched", obj, nil)

}

func AddTasks(res http.ResponseWriter, req *http.Request) {

	var obj repoTask.Task
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

	insertedId, err := behaviourTask.NewTaskRepository().AddTask(obj)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not added", nil, err.Error())
		return
	}
	if _, ok := insertedId.(int64); !ok {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not added", nil, nil)
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusCreated, "Data successfully added ", map[string]interface{}{
		"taskId": insertedId}, nil)

}

func UpdateTasksByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("taskId"))
	if err != nil || id == 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Id not provided", nil, err.Error())
		return
	}

	var reqPlayload struct {
		Name        *string     `json:"name"`
		Description interface{} `json:"description" validate:"string"`
	}

	err = helpers.BindingPayload(req, &reqPlayload)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data malformed", nil, err.Error())
		return
	}

	if reqPlayload.Name != nil && *reqPlayload.Name == "" {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data malformed", nil, "task name cannot be empty")
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
	err = behaviourTask.NewTaskRepository().UpdateTask(reqData)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not updated", nil, err.Error())
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data successfully added ", nil, nil)

}

func DeleteTasksByID(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("taskId"))
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Id not provided", nil, err.Error())
		return
	}

	if id <= 0 {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Valid Id needed", nil, errors.New("ID provided is not valid").Error())
		return
	}

	err = behaviourTask.NewTaskRepository().DeleteTask(id)
	if err != nil {
		responses.ApiResponseGivenMsg(res, http.StatusBadRequest, "Data not deleted", nil, err.Error())
		return
	}

	responses.ApiResponseGivenMsg(res, http.StatusOK, "Data successfully deleted ", nil, nil)
}
