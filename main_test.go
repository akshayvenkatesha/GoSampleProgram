package main

import (
	"Business-House-Game/constants"
	"strconv"
	"strings"
	"testing"
)

func TestInitializePlayers(t *testing.T) {
	numberOfPlayers := 3
	players := InitializePlayers(numberOfPlayers)

	for i, player := range players {
		if !strings.EqualFold(player.PlayerName, "player-"+strconv.Itoa(i+1)) {
			t.Errorf("Expected name %s but found %s", player.PlayerName, "player-"+strconv.Itoa(i))
		}

		if player.NetWorth != constants.InitialMoney {
			t.Errorf("Expected NetWorth %d but found %d", player.NetWorth, constants.InitialMoney)
		}
	}

}
