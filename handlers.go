package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ecommerce, welcome!\n")
}

func ApiItemIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(items); err != nil {
		panic(err)
	}
}

func ApiItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var itemId int
	var err error
	if itemId, err = strconv.Atoi(vars["itemId"]); err != nil {
		panic(err)
	}
	item := RepoFindItem(itemId)
	if item.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(item); err != nil {
			panic(err)
		}
		return
	}
	
	// If we didn't find it, return 404 error
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text:"Not Found"}); err != nil {
		panic(err)
	}
}

func ApiItemCreate(w http.ResponseWriter, r *http.Request) {
	var item Item
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &item); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	i := RepoCreateItem(item)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(i); err != nil{
		panic(err)
	}
}
