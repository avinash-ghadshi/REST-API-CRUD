package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	response "grad/libraries/response"

	"github.com/gorilla/mux"
)

type contact struct {
	Email   string `json:"emailid"`
	Mobile  int    `json:"mobileNumber"`
	Address string `json:"address"`
}

type UserData struct {
	Name       string  `json:"FullName"`
	Age        int     `json:"age"`
	Contacts   contact `json:"ContactDetails"`
	Experience string  `json:"exp"`
	UUID       string  `json:"aadharNumber"`
	DOB        string  `json:"BirthDate"`
}

var info map[string]UserData = map[string]UserData{}

const WelcomeMessage = `
<br><br>
<p style="text-align: center;"><strong>REST API CRUD</strong></p><br>
<p>Welcome to the simple REST API CRUD application [GOLANG].</p>
<p><strong>Available URLs:</strong></p>
<ul style="list-style-type: square;">
<li style="text-align: left;">http://127.0.0.1:8888/</li>
<li style="text-align: left;">http://127.0.0.1:8888/list/</li>
<li style="text-align: left;">http://127.0.0.1:8888/retrieve/{id}</li>
<li style="text-align: left;">http://127.0.0.1:8888/add/</li>
<li style="text-align: left;">http://127.0.0.1:8888/delete/{id}</li>
<li style="text-align: left;">http://127.0.0.1:8888/update/</li>
</ul>
<br/>
<b>Note: </b><i>{id} refers to the aadharNumber in this application. Do not put bracket in URL.<br>E.g. http://127.0.0.1:8888/retrieve/123456781290</i>
`

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(WelcomeMessage))
}

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pd UserData
	json.NewDecoder(r.Body).Decode(&pd)
	re := regexp.MustCompile(`^\d{12}$`)

	if !re.MatchString(pd.UUID) {
		response.SendResponse(w, http.StatusBadRequest, fmt.Sprintf(response.Messages["NotUUID"], pd.UUID), false)
	} else if _, ok := info[pd.UUID]; ok {
		response.SendResponse(w, http.StatusConflict, fmt.Sprintf(response.Messages["Exists"], pd.UUID), false)
	} else {
		info[pd.UUID] = pd
		response.SendResponse(w, http.StatusCreated, fmt.Sprintf(response.Messages["Success"], pd.UUID), true)
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func retrieve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	uuid := vars["UUID"]
	if data, found := info[uuid]; found {
		json.NewEncoder(w).Encode(data)
	} else {
		response.SendResponse(w, http.StatusNotFound, response.Messages["NoData"], false)
	}
}

func deleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	uuid := vars["UUID"]
	if _, found := info[uuid]; !found {
		response.SendResponse(w, http.StatusBadRequest, fmt.Sprintf(response.Messages["NotUUID"], uuid), false)
	} else {
		delete(info, uuid)
		response.SendResponse(w, http.StatusOK, fmt.Sprintf(response.Messages["Deleted"], uuid), true)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pd UserData
	json.NewDecoder(r.Body).Decode(&pd)
	re := regexp.MustCompile(`^\d{12}$`)

	if !re.MatchString(pd.UUID) {
		response.SendResponse(w, http.StatusBadRequest, fmt.Sprintf(response.Messages["NotUUID"], pd.UUID), false)
	} else if _, ok := info[pd.UUID]; ok {
		info[pd.UUID] = pd
		response.SendResponse(w, http.StatusConflict, fmt.Sprintf(response.Messages["Updated"], pd.UUID), true)
	} else {
		response.SendResponse(w, http.StatusNotFound, response.Messages["NoData"], false)
	}
}

func main() {
	router := setRoutes()
	http.ListenAndServe(":8888", router)
}
