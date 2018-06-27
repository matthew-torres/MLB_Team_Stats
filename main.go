package main

import (
	"fmt"
	"os"
)

var (
	cubs    = make(map[string]string)
	marlins = make(map[string]string)
)

func main() {
	team := os.Args[1:]

	// Chicago Cubs
	cubs["stadium"] = "Wrigley"
	cubs["manager"] = "Maddon"

	// Miami Marlins
	marlins["stadium"] = "Marlins Park"
	marlins["manager"] = "Mattingly"

	// printTeamStats using the getTeam function
	printTeamStats(getTeam(team[0]))

}

// getTeam calls a specific team based on request
func getTeam(t string) map[string]string {

	var foo map[string]string

	// Call getTeamData based on applicable team
	switch t {
	case "cubs":
		foo = cubs
	case "marlins":
		foo = marlins
	}
	return foo
}

// printTeamStats based on specific team, function prints stats for that team 
func printTeamStats(t map[string]string) {
	//Loop through team
	for k, v := range t {
		if k == "stadium" {

			//Print Stadium
			fmt.Println("Stadium: ", v)
		} else {

			//Print Manager
			fmt.Println("Manager: ", v)
		}

	}
}
