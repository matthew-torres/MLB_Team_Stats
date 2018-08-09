package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Struct used to structure the API response.
type Response struct {
	Success bool
	Message string
	Data    json.RawMessage
}

var (
	db            *sql.DB
	mysqlHost     string = "localhost"
	mysqlDatabase string = "mlb_teams_stats"
	mysqlUser     string = "root"
	mysqlPassword string = "someRandomPassword"
	dbInfo        string = mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":3306)/" + mysqlDatabase + "?charset=utf8"
	// Random API response vars
	data         json.RawMessage
	foobar       string = `{"Success": false,"Message": "Internal server error :(","Data": {"foo": "bar"}}`
	success      bool   = false
	responseCode int    = 500
	message      string
)

func main() {

	var err error

	// Flags
	team := flag.String("team", "", "A string that is used to specify a team that is used to pull data from")
	attribute := flag.String("attribute", "", "A string that is used to print out a specific stat from the team")
	flag.Parse()

	// Connect to database
	db, err = sql.Open("mysql", dbInfo)
	if err != nil {
		fmt.Printf("ERR: %s - %q", "Cannot connect to database", err)
	}

	// Get the team(s) data
	teams, err := getTeamCli(*team)
	if err != nil {
		fmt.Println(err)
	}

	if *attribute != "" {
		printTeamStats(teams, *attribute)
	} else {
		printTeamStats(teams, "")
	}

	// Start the mux router
	r := Router()

	// Launch the server
	log.Fatal(http.ListenAndServe(":8080", r))

	defer db.Close()

}

// Router is used to start the api router, and define the endpoint which API
// consumers can route to.
func Router() *mux.Router {

	// Start the router
	r := mux.NewRouter()

	// Team endpoints
	r.HandleFunc("/api/v1/team", AddTeam).Methods("PUT")
	r.HandleFunc("/api/v1/team/{team}", GetTeam).Methods("GET")

	// Player endpoints
	r.HandleFunc("/api/v1/player", AddPlayer).Methods("PUT")
	r.HandleFunc("/api/v1/player/{player}", GetPlayer).Methods("GET")

	return r
}
