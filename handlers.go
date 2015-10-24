package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
func LocationShow(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
    todoId := vars["location_id"]
		l, err := strconv.Atoi(todoId)
		if err != nil {
			panic(err)
		}

		var getLoc Location
		getLoc=RepoShowLocation(l)
    //fmt.Fprintln(w, "Todo show:", getLoc)
		res2B, _ :=json.Marshal(getLoc)
		w.Header().Set("Content-Type", "application/json;")
		w.WriteHeader(200)
		fmt.Fprintf(w,"%s",res2B)

}
func LocationUpdate(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	todoId := vars["location_id"]
	l, err := strconv.Atoi(todoId)
	if err != nil {
		panic(err)
	}
	var loc Location
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &loc); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		/*if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}*/
	}
	Nc:=QueryGMaps(loc)

resp := Location{
	Id:loc.Id,
	Name:loc.Name,
	Address:loc.Address,
	City:loc.City,
	State:loc.State,
	Zip:loc.Zip,
	Coordinate:Nc,
}
mess:=	RepoUpdateLocation(l,resp)

	//json.NewEncoder(w).Encode(mess)
	res2B, _ :=json.Marshal(mess)
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w,"%s",res2B)

}
func LocationRemove(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	todoId := vars["location_id"]
	l, err := strconv.Atoi(todoId)
	if err != nil {
		panic(err)
	}

	RepoRemoveLocation(l)
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(200)

}
func LocationCreate(rw http.ResponseWriter, r *http.Request) {
	//t := RepoCreateTodo(todo)
	var loc Location

  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &loc); err != nil {
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//rw.WriteHeader(422) // unprocessable entity
		/*if err := json.NewEncoder(rw).Encode(err); err != nil {
			panic(err)
		}*/
	}
	Nc:=QueryGMaps(loc)
	var mess msg

resp := Location{
	Id:mess.Id,
	Name:loc.Name,
	Address:loc.Address,
	City:loc.City,
	State:loc.State,
	Zip:loc.Zip,
	Coordinate:Nc,
}
//js, _ := json.Marshal(resp)
//fmt.Printf("%s", js)
mess=RepoAddLocation(resp)
respnew := Location{
	Id:mess.Id,
	Name:loc.Name,
	Address:loc.Address,
	City:loc.City,
	State:loc.State,
	Zip:loc.Zip,
	Coordinate:Nc,
}
res2B, _ :=json.Marshal(respnew)
rw.Header().Set("Content-Type", "application/json;")
rw.WriteHeader(http.StatusCreated)
fmt.Fprintf(rw,"%s",res2B)

}
