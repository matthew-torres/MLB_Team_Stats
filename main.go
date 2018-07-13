package main

import (
	"fmt"
	"os"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type Team struct {
	ID int `json:"id" sql:"id"`
	Name string `json:"name" sql:"name"`
	Manager string `json:"manager" sql:"manager"`
	Stadium string `json:"stadium" sql:"stadium"`
	City string `json:"city" sql:"city"`
	State string `json:"state" sql:"state"`
}

var (
	cubs    = make(map[string]string)
	marlins = make(map[string]string)
	db            *sql.DB
	mysqlHost     string = "localhost"
	mysqlDatabase string = "mlb_team_stats"
	mysqlUser     string = "root"
	mysqlPassword string = "pa11word"
	dbInfo        string = mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":3306)/" + mysqlDatabase + "?charset=utf8"
)

func main() {
	team := os.Args[1:]

	// Connect to database
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		fmt.Printf("ERR: %s - %q", "Cannot connect to database", err)
	}
	fmt.Println(team)
	defer db.Close()

}

// getTeam calls a specific team based on request
func getTeam(t string) (teams []Team,err error) {

	var rows *sql.Rows

	// Get all assets
	rows, err = db.Query("SELECT * FROM teams WHERE name = (?)",t)

	// Check to make sure there were no errors in querying the data
	if err != nil {
		// Could not query the db, send empty response and error.
		fmt.Printf("ERROR: %q",err)
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
			fmt.Printf("ERROR: %q",err)
		}
		teams = append(teams, t)
	}
	return teams,nil
}

// printTeamStats based on specific team, function prints stats for that team
func printTeamStats(t map[string]string,a string) {

	// Check if a specific attribute was requested
	if a == "" {
		//Loop through team
		for k, v := range t {
			fmt.Println(strings.Title(k), ":", v)
		}
	} else {
		// Check if the requested attribute exists
		if t[a] == "" {
			fmt.Println("Invalid attribute requested")
		} else {
			fmt.Println(t[a])
		}
	}
}
