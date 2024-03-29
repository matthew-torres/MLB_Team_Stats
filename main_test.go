package main

import (
	"testing"
)

// TestGetTeam test to see if the right value is retrieved
func TestGetTeam(t *testing.T) {

	// Chicago Cubs
	cubs["stadium"] = "Wrigley"
	cubs["manager"] = "Maddon"

	// Miami Marlins
	marlins["stadium"] = "Marlins Park"
	marlins["manager"] = "Mattingly"

	bar, err := getTeam("cubs")
	if err != nil {
		t.Errorf("Received invalid team: %q", err)
	}
	if bar["stadium"] != cubs["stadium"] {
		t.Errorf("Expecting %s got %s", cubs["stadium"], bar["stadium"])
	}
}
