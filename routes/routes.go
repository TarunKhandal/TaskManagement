package routes

import (
	"TaskManagement/app/http/controllers"
	middlewares "TaskManagement/app/http/middlerwares"
	"encoding/json"
	"net/http"
)

func Routes() http.Handler {

	// new router instance
	mainGrp := http.NewServeMux()

	mainGrp.Handle("/", middlewares.RcoverMiddlerware(taskRoutes()))

	return mainGrp

}

func taskRoutes() http.Handler {
	// new router instance
	tr := http.NewServeMux()

	// Authors Routes
	tr.Handle("GET /author", http.HandlerFunc(controllers.GetAllAuthor))
	tr.Handle("GET /author/{authorId}", http.HandlerFunc(controllers.GetAuthorByID))
	tr.Handle("POST /author", http.HandlerFunc(controllers.AddAuthor))
	tr.Handle("PUT /author/{authorId}", http.HandlerFunc(controllers.UpdateAuthorByID))
	tr.Handle("DELETE /author/{authorId}", http.HandlerFunc(controllers.DeleteAuthorByID))

	// Tasks Routes
	tr.Handle("GET /tasks", http.HandlerFunc(controllers.GetAllTasks))
	tr.Handle("GET /tasks/{authorId}/task", http.HandlerFunc(controllers.GetTasksByAuthorID))
	tr.Handle("GET /tasks/{taskId}", http.HandlerFunc(controllers.GetTasksByID))
	tr.Handle("POST /tasks", http.HandlerFunc(controllers.AddTasks))
	tr.Handle("PUT /tasks/{taskId}", http.HandlerFunc(controllers.UpdateTasksByID))
	tr.Handle("DELETE /tasks/{taskId}", http.HandlerFunc(controllers.DeleteTasksByID))

	// Steps Routes
	tr.Handle("GET /tasks/{taskId}/steps", http.HandlerFunc(controllers.GetStepsByTaskID))
	tr.Handle("GET /tasks/{taskId}/steps/{stepId}", http.HandlerFunc(controllers.GetStepByTaskID))
	tr.Handle("POST /tasks/{taskId}/steps", http.HandlerFunc(controllers.AddStepsByTaskID))
	tr.Handle("PUT /tasks/{taskId}/steps/{stepId}", http.HandlerFunc(controllers.UpdateStepsByID))
	tr.Handle("DELETE /tasks/{taskId}/steps/{stepId}", http.HandlerFunc(controllers.DeleteStepsByID))

	tr.Handle("/", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		respose := map[string]string{
			"error": "This Method is not allowed for this api",
		}
		writeOut, _ := json.Marshal(respose)
		res.Write(writeOut)
	}))

	return tr
}
