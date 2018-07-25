package main

import (
	"fmt"
	"flag"
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

	var err error

	// Flags
	team := flag.String("team","", "A string that is used to specify a team that is used to pull data from")
	attribute := flag.String("attribute","", "A string that is used to print out a specific stat from the team")
	flag.Parse()

	// Connect to database
	db, err = sql.Open("mysql", dbInfo)
	if err != nil {
		fmt.Printf("ERR: %s - %q", "Cannot connect to database", err)
	}

	// Get the team(s) data
	teams,err := getTeam(*team)
	if err != nil {
		fmt.Println(err)
	}

	if *attribute != "" {
		printTeamStats(teams,*attribute)
	} else {
		printTeamStats(teams,"")
	}

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
//func printTeamStats(t map[string]string,a string) {
func printTeamStats(teams []Team,a string) {

	// Check if a specific attribute was requested
	if a == "" {
		// Print all attributes
		for _,team := range teams {

			fmt.Println(team.ID)
			fmt.Println(team.Name)
			fmt.Println(team.Manager)
			fmt.Println(team.Stadium)
			fmt.Println(team.City)
			fmt.Println(team.State)
		}

	} else {

		//Loop through team and print specific attribute
		for _,team := range teams {
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
