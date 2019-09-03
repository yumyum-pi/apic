package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yumyum-pi/apic/model"
	"github.com/yumyum-pi/apic/model/db"
	"github.com/yumyum-pi/apic/route/utility"
)

var database *sql.DB

// get users
func getOneUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, "Invalid users ID")
		return
	}
	u := model.Users{ID: id}
	if err := u.GetUsers(database); err != nil {
		switch err {
		case sql.ErrNoRows:
			utility.RespondWithError(w, http.StatusNotFound, "Users not found")
		default:
			utility.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utility.RespondWithJSON(w, http.StatusOK, u)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}
	users, err := model.GetUsers(database, start, count)
	if err != nil {
		utility.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utility.RespondWithJSON(w, http.StatusOK, users)
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	var u model.Users
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := u.CreateUsers(database); err != nil {
		utility.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utility.RespondWithJSON(w, http.StatusCreated, u)
}

func updateUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, "Invalid users ID")
		return
	}
	var u model.Users
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	u.ID = id
	if err := u.UpdateUsers(database); err != nil {
		utility.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utility.RespondWithJSON(w, http.StatusOK, u)
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, "Invalid Users ID")
		return
	}
	u := model.Users{ID: id}
	if err := u.DeleteUsers(database); err != nil {
		utility.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utility.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func deleteManyUsers(w http.ResponseWriter, r *http.Request) {
	var u []model.Users
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if len(u) > 10 {
		utility.RespondWithError(w, http.StatusBadRequest, "Exceed Max ID per request. Max ID per request = 10")
	} else if len(u) < 0 {
		utility.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	if err := model.DeleteUsers(database, u); err != nil {
		utility.RespondWithError(w, http.StatusBadRequest, err.Error())
	}

	utility.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// UsersRouter handles all the employrouter
func UsersRouter(r *mux.Router) {
	r.HandleFunc("/", getUsers).Methods("GET")
	r.HandleFunc("/", createUsers).Methods("POST")
	r.HandleFunc("/", deleteManyUsers).Methods("DELETE")
	r.HandleFunc("/{id:[0-9]+}", getOneUsers).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", updateUsers).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}", deleteUsers).Methods("DELETE")
}

// UsersRouteInit interlistion the database variable
func UsersRouteInit() {
	database = db.GetDB()
}
