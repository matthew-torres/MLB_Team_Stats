package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Team struct {
	ID      int    `json:"id" sql:"id"`
	Name    string `json:"name" sql:"name"`
	Manager string `json:"manager" sql:"manager"`
	Stadium string `json:"stadium" sql:"stadium"`
	City    string `json:"city" sql:"city"`
	State   string `json:"state" sql:"state"`
}

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

// getTeam calls a specific team based on request
func getTeamCli(t string) (teams []Team, err error) {

	var rows *sql.Rows

	// Get all assets
	rows, err = db.Query("SELECT * FROM teams WHERE name = (?)", t)

	// Check to make sure there were no errors in querying the data
	if err != nil {
		// Could not query the db, send empty response and error.
		fmt.Printf("ERROR: %q", err)
	}

	// Defer the cosing of the db connection
	defer rows.Close()

	// Iterate through the rows and add them to the assets.
	for rows.Next() {
		t := Team{}
		// Scan the results into the assets variable type Assets
		err := rows.Scan(&t.ID, &t.Name, &t.Manager, &t.Stadium, &t.City, &t.State)

		// Check for error during scan
		if err != nil {
			// Could not scan query results into assets, retun empty van and error
			fmt.Printf("ERROR: %q", err)
		}
		teams = append(teams, t)
	}
	return teams, nil
}

// printTeamStats based on specific team, function prints stats for that team
//func printTeamStats(t map[string]string,a string) {
func printTeamStats(teams []Team, a string) {

	// Check if a specific attribute was requested
	if a == "" {
		// Print all attributes
		for _, team := range teams {

			fmt.Println(team.ID)
			fmt.Println(team.Name)
			fmt.Println(team.Manager)
			fmt.Println(team.Stadium)
			fmt.Println(team.City)
			fmt.Println(team.State)
		}

	} else {

		//Loop through team and print specific attribute
		for _, team := range teams {
			switch a {
			case "id":
				fmt.Println(team.ID)
			case "name":
				fmt.Println(team.Name)
			case "manager":
				fmt.Println(team.Manager)
			case "stadium":
				fmt.Println(team.Stadium)
			case "city":
				fmt.Println(team.City)
			case "state":
				fmt.Println(team.State)
			default:
				fmt.Println("Attribute not found")
			}

		}
	}
}

// Router is used to start the api router, and define the endpoint which API
// consumers can route to.
func Router() *mux.Router {

	// Start the router
	r := mux.NewRouter()

	// Assets defined in logs.go
	r.HandleFunc("/api/v1/team", AddTeam).Methods("PUT")
	//r.HandleFunc("/api/v1/teams", GetTeams).Methods("GET")
	r.HandleFunc("/api/v1/team/{team}", GetTeam).Methods("GET")

	return r
}

func GetTeam(w http.ResponseWriter, r *http.Request) {

	// Get the team name from the request
	vars := mux.Vars(r)
	team := vars["team"]

	// Now that we have the name from the API request, query the database for the requested team

	t, err := getTeamCli(team)
	if err != nil {
		log.Println("ERROR: Cannot find requested team - %q", err)
	} else {

		// Check if team exists
		if len(t) > 0 {

			// Team exists, format response
			io.WriteString(w, t[0].Name)

			// @TODO - proper JSON responsei - START HERE

		} else {

			// Respond with applicable message and JSON
			io.WriteString(w, "No data found")

		}
	}

}

func AddTeam(w http.ResponseWriter, r *http.Request) {

	// The team payload
	p := json.NewDecoder(r.Body)

	var team Team

	// Decode the payload in to 'team' [Team]
	err := p.Decode(&team)
	if err != nil {
		success = false
		responseCode = 500
		message = "Internal Error :("
		log.Printf("ERR: Could not decode payload - %q", err)
	} else {

		// Query DB to check for existing team

		// Team already exists
		if 1 == 2 {
			// Output message informing user that team already exists in DB
		} else {
			// Insert data into DB
			_, err := db.Exec("INSERT INTO teams (name,manager,stadium,city,state) values (?,?,?,?,?)", team.Name, team.Manager, team.Stadium, team.City, team.State)
			if err != nil {
				success = false
				responseCode = 500
				message = "Internal Error :("
				log.Printf("ERR: Could not decode payload - %q", err)
			} else {

				// All Good
				success = true
				responseCode = 202 // Accepted
				message = "Request accepted, Team added."
			}
		}

	}

	// By this point we should have some sort of response
	resp := &Response{Success: success, Message: message, Data: data}

	// SET content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Marshal the response
	response, err := json.Marshal(resp)

	// Check to see if there was an error whilst marshalling the response
	if err != nil {

		// FML
		log.Printf("ERR: Could not marshal response - %q", err)
		w.WriteHeader(500)
		fmt.Fprint(w, foobar)

	} else {

		// Respond
		w.WriteHeader(responseCode)
		fmt.Fprint(w, string(response))
	}

}
