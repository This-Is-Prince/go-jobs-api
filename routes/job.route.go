package routes

import (
	"net/http"

	"github.com/This-Is-Prince/go-jobs-api/controllers"
	"github.com/gorilla/mux"
)

func JobRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		if method == "POST" {
			controllers.CreateJob(w, r)

		} else if method == "GET" {
			controllers.GetAllJobs(w, r)
		} else {
			controllers.DeleteAllJobs(w, r)
		}
	}).Methods("POST", "GET", "DELETE")

	return r
}
