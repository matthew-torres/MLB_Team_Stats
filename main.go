package main

import (
	"fmt"
	"os"
)

func main() {
	team := os.Args[1:]
	cubs := make(map[string]string)
	cubs["stadium"] = "Wrigley"
	cubs["manager"] = "Madodon"
	getTeamData(team)
}

func getTeamData(t []string) {
	//Loop through team
		//Print Manager
		//Print Stadium 
}
