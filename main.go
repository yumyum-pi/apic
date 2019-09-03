package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/yumyum-pi/apic/model/db"
	rs "github.com/yumyum-pi/apic/route"
)

var ip string
var port string

// Init Initialization of the API sever
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ip = os.Getenv("API_IP")
	port = os.Getenv("API_PORT")
}
func welcomeToAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the api")
}

func main() {
	Init()

	r := mux.NewRouter()
	r.HandleFunc("/", welcomeToAPI).Methods("GET")

	// Initialization of Database
	db.Init()

	//Initialization and adding route for users
	rs.UsersRouteInit()
	rs.UsersRouter(r.PathPrefix("/api/users").Subrouter())

	/*route*/

	// Start server
	addr := fmt.Sprintf("%s:%s", ip, port)
	fmt.Printf(" > starting server @ %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
