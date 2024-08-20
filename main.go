package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/iancoleman/strcase"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

var DB *sqlx.DB

func main() {

	DB = InitDB()

	r := mux.NewRouter()
	r.Use(headerMiddleware, requestMiddleware)

	// Handle all preflight request
	// todo: are all methods needed?
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	// Endpoints
	endpoints := [5]string{
		"/",
		"/nodes",
		"/nodes/{id:[0-9]+}",
		"/discussion",
		"/discussion/{id:[0-9]+}",
	}

	// Route Handlers
	r.HandleFunc(endpoints[0], About).Methods("GET")
	r.HandleFunc(endpoints[1], ReadNodes).Methods("GET")
	r.HandleFunc(endpoints[1], CreateNode).Methods("POST")
	//r.HandleFunc(endpointInteger, ReadItemRecord).Methods("GET")
	//r.HandleFunc(endpointInteger, UpdateItemRecord).Methods("PUT")
	//r.HandleFunc(endpointInteger, DeleteItemRecord).Methods("DELETE")

	r.NotFoundHandler = http.HandlerFunc(HTTPNotFound)

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func requestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Incoming request ", r.Method)
		next.ServeHTTP(w, r)
	})
}

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:password@tcp(localhost:3306)/atm") // "jdbc:sqlite:identifier.sqlite")
	HandleError(err)
	return db
}
