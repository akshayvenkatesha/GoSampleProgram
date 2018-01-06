package main

import (
	"Business-House-Game/constants"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestCalculateMain(t *testing.T) {
	NumberOfPlayers := 3
	cellsLayout := []string{"E", "E", "J", "H", "E", "T", "J", "T", "E", "E", "H", "J", "T", "H", "E", "E", "J", "H", "E", "T", "J", "T", "E", "E", "H", "J", "T", "H", "J", "E", "E", "J", "H", "E", "T", "J", "T", "E", "E", "H", "J", "T", "E", "H", "E"}
	diceOutArray := []int{4, 4, 4, 6, 7, 8, 5, 11, 10, 12, 2, 3, 5, 6, 7, 8, 5, 11, 10, 12, 2, 3, 5, 6, 7, 8, 5, 11, 10, 12}
	playerPairs := InitializePlayers(NumberOfPlayers)

	finalOutput := StartGame(playerPairs, cellsLayout, diceOutArray)

	for _, player := range finalOutput {
		fmt.Printf("%d has total worth %d", player.PlayerName, player.NetWorth)
	}
}

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
