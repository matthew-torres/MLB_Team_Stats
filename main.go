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

	// Call getTeamData based on applicable team
	switch team[0] {
	case "cubs":
		getTeamData(cubs)
	case "marlins":
		getTeamData(marlins)
	}

}

func getTeamData(t map[string]string) {
	//Loop through team
	for k, v := range t {
		if k == "stadium" {

			//Print Stadium
			fmt.Println("stadium: ", v)
		} else {

			//Print Manager
			fmt.Println("manager: ", v)
		}
	}
}
