package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Player struct {
	ID        int    `json:"id" sql:"id"`
	Team_id   string `json:"team_id" sql:"team_id"`
	Firstname string `json:"firstname" sql:"firstname"`
	Lastname  string `json:"lastname" sql:"lastname"`
	Position  string `json:"position" sql:"position"`
	Number    int    `json:"number" sql:"number"`
}

func AddPlayer(w http.ResponseWriter, r *http.Request) {

	log.Println("oof")
	log.Println(r.Body)

	// The player payload
	p := json.NewDecoder(r.Body)

	var player Player

	// Decode the payload in to 'player' [Player]
	err := p.Decode(&player)
	if err != nil {
		success = false
		responseCode = 500
		message = "Internal Error :("
		log.Printf("ERR: Could not decode payload - %q", err)
	} else {

		// Query DB to check for existing player

		// Player already exists
		if 1 == 2 {
			// Output message informing user that player already exists in DB
		} else {
			// Insert data into DB
			_, err := db.Exec("INSERT INTO players (team_id,firstname,lastname,position,number) values (?,?,?,?,?)", player.Team_id, player.Firstname, player.Lastname, player.Position, player.Number)
			if err != nil {
				success = false
				responseCode = 500
				message = "Internal Error :("
				log.Printf("ERR: Could not decode payload - %q", err)
			} else {

				// All Good
				success = true
				responseCode = 202 // Accepted
				message = "Request accepted, Player added."
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

func GetPlayer(w http.ResponseWriter, r *http.Request) {

	// Get the player name from the request
	vars := mux.Vars(r)
	player := vars["team"]

	// Now that we have the name from the API request, query the database for the requested player

	t, err := getPlayerCli(player)
	if err != nil {
		log.Println("ERROR: Cannot find requested player - %q", err)
	} else {

		// Check if player exists
		if len(t) > 0 {

			// Player exists, format response
			io.WriteString(w, t[0].Lastname)

			// @TODO - proper JSON responsei - START HERE

		} else {

			// Respond with applicable message and JSON
			io.WriteString(w, "No data found")

		}
	}

}

// getPlayer calls a specific player based on request
func getPlayerCli(p string) (players []Player, err error) {

	var rows *sql.Rows

	// Get all assets
	rows, err = db.Query("SELECT * FROM players WHERE name = (?)", p)

	// Check to make sure there were no errors in querying the data
	if err != nil {
		// Could not query the db, send empty response and error.
		fmt.Printf("ERROR: %q", err)
	}

	// Defer the cosing of the db connection
	defer rows.Close()

	// Iterate through the rows and add them to the assets.
	for rows.Next() {
		p := Player{}
		// Scan the results into the assets variable type Assets
		err := rows.Scan(&p.ID, &p.Team_id, &p.Firstname, &p.Lastname, &p.Position, &p.Number)

		// Check for error during scan
		if err != nil {
			// Could not scan query results into assets, retun empty van and error
			fmt.Printf("ERROR: %q", err)
		}
		players = append(players, p)
	}
	return players, nil
}
