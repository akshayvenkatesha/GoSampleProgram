package main

import (
	"Business-House-Game/constants"
	"Business-House-Game/playerPairArray"
	"Business-House-Game/readInputs"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var NumberOfPlayers int
	var CellsLayout []string
	var diceOutArray []int
	var playerPairs playerPairArray.PlayerPairArray

	NumberOfPlayers, errPlayerRead := readInputs.ReadNumberOfPlayers()
	if errPlayerRead != nil {
		fmt.Println(errPlayerRead)
		return
	}

	CellsLayout, errCellRead := readInputs.ReadCellsPositonAndTypes()
	if errCellRead != nil {
		fmt.Println(errCellRead)
		return
	}

	diceOutArray, errDiceRead := readInputs.ReadDiceValues()
	if errDiceRead != nil {
		fmt.Println(errDiceRead)
		return
	}

	playerPairs = InitializePlayers(NumberOfPlayers)

	finalOutput := StartGame(playerPairs, CellsLayout, diceOutArray)

	for _, player := range finalOutput {
		fmt.Printf("%d has total worth %d", player.PlayerName, player.NetWorth)
	}
}

func StartGame(playersPairs playerPairArray.PlayerPairArray, cellsLayout []string, diceOutArray []int) playerPairArray.PlayerPairArray {

	hotelOwerDict := make(map[int]int)
	i := 0
	for _, diceValue := range diceOutArray {
		playersPairs[i].CurrentPosition += diceValue

		for playersPairs[i].CurrentPosition >= len(cellsLayout) {
			playersPairs[i].CurrentPosition = playersPairs[i].CurrentPosition % len(cellsLayout)
		}

		if strings.EqualFold(cellsLayout[playersPairs[i].CurrentPosition], constants.Jail) {
			playersPairs[i].NetWorth = playersPairs[i].NetWorth - constants.JailFine
		}

		if strings.EqualFold(cellsLayout[playersPairs[i].CurrentPosition], constants.Treasure) {
			playersPairs[i].NetWorth = playersPairs[i].NetWorth + constants.TreasureValue
		}

		if strings.EqualFold(cellsLayout[playersPairs[i].CurrentPosition], constants.Hotel) {
			ownerIndex, owned := hotelOwerDict[playersPairs[i].CurrentPosition]

			if owned && ownerIndex != playersPairs[i].PlayerIndex {
				playersPairs[i].NetWorth = playersPairs[i].NetWorth - constants.HotelRent
				playersPairs[ownerIndex].NetWorth = playersPairs[ownerIndex].NetWorth + constants.HotelRent
			}
			if !owned && playersPairs[i].NetWorth >= constants.HotelWorth {
				playersPairs[i].NetWorth = playersPairs[i].NetWorth - constants.HotelWorth
				hotelOwerDict[playersPairs[i].CurrentPosition] = playersPairs[i].PlayerIndex
			}
		}

		i++
		if i == len(playersPairs) {
			i = i % len(playersPairs)
		}
	}

	for _, playerIndex := range hotelOwerDict {
		playersPairs[playerIndex].NetWorth = playersPairs[playerIndex].NetWorth + constants.HotelWorth
	}

	return playerPairArray.Sort(playersPairs)

}

func InitializePlayers(NumberOfPlayers int) playerPairArray.PlayerPairArray {
	var playerArray playerPairArray.PlayerPairArray

	for i := 1; i <= NumberOfPlayers; i++ {
		playerPair := playerPairArray.PlayerPair{
			PlayerIndex: i - 1,
			PlayerName:  "Player-" + strconv.Itoa(i),
			NetWorth:    constants.InitialMoney, CurrentPosition: 0}
		playerArray = append(playerArray, playerPair)
	}
	return playerArray
}
