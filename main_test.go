package main

import (
	"Business-House-Game/constants"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestValidateStartGameAllEmpty(t *testing.T) {
	NumberOfPlayers := 3
	cellsLayout := []string{"E", "E", "E", "E", "E", "E", "E"}
	diceOutArray := []int{4, 4, 4, 6, 7, 8, 5, 11, 10, 12, 2, 3, 5, 6, 7, 8, 5, 11, 10, 12, 2, 3, 5, 6, 7, 8, 5, 11, 10, 12}
	playerPairs := InitializePlayers(NumberOfPlayers)

	finalOutput := StartGame(playerPairs, cellsLayout, diceOutArray)

	for _, player := range finalOutput {
		if player.NetWorth != constants.InitialMoney {
			t.Errorf("Expected %d for playter %s but found %d", constants.InitialMoney, player.PlayerName, player.NetWorth)
		}
	}
}

func TestValidateStartGameAllJail(t *testing.T) {
	NumberOfPlayers := 3
	cellsLayout := []string{"E", "J", "J", "J", "J"}
	diceOutArray := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	playerPairs := InitializePlayers(NumberOfPlayers)

	finalOutput := StartGame(playerPairs, cellsLayout, diceOutArray)

	for _, player := range finalOutput {
		if player.NetWorth != (constants.InitialMoney - constants.JailFine*(len(cellsLayout)-1)) {
			t.Errorf("Expected %d for playter %s but found %d", (constants.InitialMoney - (constants.JailFine*len(cellsLayout) - 1)), player.PlayerName, player.NetWorth)
		}
		if player.CurrentPosition != 0 {
			t.Errorf("Expected %d for current postion for playter %s but found %d", 0, player.PlayerName, player.CurrentPosition)
		}
	}
}

func TestValidateStartGameAllTreasures(t *testing.T) {
	NumberOfPlayers := 3
	cellsLayout := []string{"E", "T", "T", "T", "T"}
	diceOutArray := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	playerPairs := InitializePlayers(NumberOfPlayers)

	finalOutput := StartGame(playerPairs, cellsLayout, diceOutArray)

	for _, player := range finalOutput {
		if player.NetWorth != (constants.InitialMoney + constants.TreasureValue*(len(cellsLayout)-1)) {
			t.Errorf("Expected %d for playter %s but found %d", (constants.InitialMoney - (constants.TreasureValue*len(cellsLayout) - 1)), player.PlayerName, player.NetWorth)
		}
		if player.CurrentPosition != 0 {
			t.Errorf("Expected %d for current postion for playter %s but found %d", 0, player.PlayerName, player.CurrentPosition)
		}
	}
}

func TestValidateStartGameAllHotels(t *testing.T) {
	NumberOfPlayers := 3
	cellsLayout := []string{"E", "H", "H", "H", "H"}
	diceOutArray := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	playerPairs := InitializePlayers(NumberOfPlayers)

	finalOutput := StartGame(playerPairs, cellsLayout, diceOutArray)

	assert.Equal(t, finalOutput[0].NetWorth, 1400)
	assert.Equal(t, finalOutput[1].NetWorth, 800)
	assert.Equal(t, finalOutput[2].NetWorth, 800)

	for _, player := range finalOutput {
		if player.CurrentPosition != 0 {
			t.Errorf("Expected %d for current postion for playter %s but found %d", 0, player.PlayerName, player.CurrentPosition)
		}
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
