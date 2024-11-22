package rest

import (
	"Hausuebung-I/repo"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Rest struct {
	Router *mux.Router
	Repo   *repo.Repo
}

func New(router *mux.Router, repo *repo.Repo) (rest Rest) {
	rest.Router = router
	rest.Repo = repo
	rest.Router.HandleFunc("/issue", rest.PostIssue).Methods("POST")
	rest.Router.HandleFunc("/issue/{id}", rest.GetIssue).Methods("GET")
	rest.Router.HandleFunc("/issue/{id}", rest.PutIssue).Methods("PUT")
	rest.Router.HandleFunc("/issue/{id}", rest.DeleteIssue).Methods("DELETE")
	rest.Router.HandleFunc("/issue/{id}", rest.PatchIssue).Methods("PATCH")
	return
}

func (rest *Rest) PostIssue(w http.ResponseWriter, r *http.Request) {
	var issue repo.Issue
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&issue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	defer r.Body.Close()

	emptyId := uuid.UUID{}
	if issue.Id == emptyId {
		issue.Id = uuid.New()
	}

	if err = (*(*rest).Repo).CreateIssue(issue); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(issue)
}

func (rest *Rest) GetIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var (
		issue repo.Issue
		id    uuid.UUID
		err   error
	)

	if id, err = uuid.Parse(vars["id"]); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	if issue, err = (*(*rest).Repo).GetIssue(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(issue)
}

func (rest *Rest) PutIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var (
		issue repo.Issue
		id    uuid.UUID
		err   error
	)
	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(&issue); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	if id, err = uuid.Parse(vars["id"]); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	if err = (*(*rest).Repo).PutIssue(id, issue); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (rest *Rest) PatchIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var (
		issue repo.Issue
		id    uuid.UUID
		err   error
	)
	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(&issue); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	if id, err = uuid.Parse(vars["id"]); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	if err = (*(*rest).Repo).PatchIssue(id, issue); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (rest *Rest) DeleteIssue(w http.ResponseWriter, r *http.Request) {
	var (
		id  uuid.UUID
		err error
	)
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	if id, err = uuid.Parse(vars["id"]); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	if err = (*(*rest).Repo).DeleteIssue(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
