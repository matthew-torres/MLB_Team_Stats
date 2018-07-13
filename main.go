package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"errors"
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
	cubs["location"] = "Chicago, Illinois"
	cubs["team average"] = strconv.Itoa(266)
	cubs["hits"] = strconv.Itoa(792)

	// Miami Marlins
	marlins["stadium"] = "Marlins Park"
	marlins["manager"] = "Mattingly"
	marlins["location"] = "Miami, Florida"
	marlins["average"] = strconv.Itoa(242)
	marlins["hits"] = strconv.Itoa(751)

	// printTeamStats using the getTeam function
	data,err := getTeam(team[0])
	if err != nil {
		fmt.Println(err)
	} else {
		// Check for second paramater
		if (len(team) == 2) { // ==2 so we have a second paramater
			printTeamStats(data,team[1])
		} else {
			printTeamStats(data,"")
		}

	}

}

// getTeam calls a specific team based on request
func getTeam(t string) (map[string]string,error) {

	var foo map[string]string

	// Call getTeamData based on applicable team
	switch t {
	case "cubs":
		foo = cubs
	case "marlins":
		foo = marlins
	}
	if len(foo) > 0 {
		return foo,nil
	} else {
		return foo,errors.New("Team not found.")
	}
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
