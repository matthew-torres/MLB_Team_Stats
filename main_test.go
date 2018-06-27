package main 

import (
	"testing"
	"fmt"
)

var (
	team1    = make(map[string]string)
	team2 = make(map[string]string)
)

func TestGetTeam(t *testing.T) {
	
	// Chicago Cubs
	team1["stadium"] = "Wrigley"
	team1["manager"] = "Maddon"

	// Miami Marlins
	team2["stadium"] = "Marlins Park"
	team2["manager"] = "Mattingly"
	
	bar := getTeam("cubs")
	fmt.Println(foo)
}

