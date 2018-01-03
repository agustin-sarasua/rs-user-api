package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	c "github.com/agustin-sarasua/rs-common"
	m "github.com/agustin-sarasua/rs-model"
	"github.com/gorilla/mux"
)

func CreateUserEndpoint(w http.ResponseWriter, req *http.Request) {
	var msg m.User
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		c.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}
	msg.CreatedAt = time.Now()
	if id, errs := CreateUser(&msg); len(errs) > 0 {
		log.Printf("Error creating user")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(m.BuildErrorResponse(errs))
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "{id: %q}", id)
	}
	w.Header().Set("Content-Type", "application/json")
}

func UpdateUserEndpoint(w http.ResponseWriter, req *http.Request) {
	var msg m.User
	err := json.NewDecoder(req.Body).Decode(&msg)

	if err != nil {
		c.ErrorWithJSON(w, "", http.StatusBadRequest)
		return
	}
	msg.CreatedAt = time.Now()
	if id, errs := CreateUser(&msg); len(errs) > 0 {
		log.Printf("Error creating user")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(m.BuildErrorResponse(errs))
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "{id: %q}", id)
	}
	w.Header().Set("Content-Type", "application/json")
}

func GetUserByFirebaseIDEndpoint(w http.ResponseWriter, req *http.Request) *m.User {
	v := req.URL.Query()
	fbID, ok := v["firebaseID"]
	var u m.User
	if ok {
		Db.First(&u, "FirebaseID = ?", fbID)
		return &u
	}
	return nil
}

func GetUserEndpoint(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.ParseUint(mux.Vars(req)["id"], 10, 64)

	u := LoadUser(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
