package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sesi9-project/controllers"
	"sesi9-project/middleware"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Printf("server started at localhost%s\n", server.Addr)
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !middleware.Auth(w, r) {
		return
	}

	if !middleware.AllowOnlyGET(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); (id) != "" {
		OutputJSON(w, controllers.SelectStudent(id))
		return
	}

	OutputJSON(w, controllers.GetStudents())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
