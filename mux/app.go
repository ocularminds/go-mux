package mux

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//App models the application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize initialises the application database credentials
func (a *App) Initialize(user, pass, dbname string) {
	con := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, dbname)

	var err error
	a.DB, err = sql.Open("postgres", con)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

//Run runs the application
func (a *App) Run(addr string) {}
