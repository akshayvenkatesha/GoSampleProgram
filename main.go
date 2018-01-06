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
	fmt.Println(NumberOfPlayers)

	CellsLayout, errCellRead := readInputs.ReadCellsPositonAndTypes()
	if errCellRead != nil {
		fmt.Println(errCellRead)
		return
	}
	fmt.Println(CellsLayout)

	diceOutArray, errDiceRead := readInputs.ReadDiceValues()
	if errDiceRead != nil {
		fmt.Println(errDiceRead)
		return
	}
	fmt.Println(diceOutArray)

	playerPairs = InitializePlayers(NumberOfPlayers)

	StartGame(playerPairs, CellsLayout, diceOutArray)
}

func StartGame(playersPairs playerPairArray.PlayerPairArray, cellsLayout []string, diceOutArray []int) {

	hotelOwerDict := make(map[int]int)
	i := 0
	for _, diceValue := range diceOutArray {
		player := playersPairs[i]
		player.CurrentPosition += diceValue

		if player.CurrentPosition > len(cellsLayout) {
			player.CurrentPosition = player.CurrentPosition % len(cellsLayout)
		}

		if strings.EqualFold(cellsLayout[player.CurrentPosition], constants.Jail) {
			player.NetWorth -= constants.JailFine
		}

		if strings.EqualFold(cellsLayout[player.CurrentPosition], constants.Treasure) {
			player.NetWorth += constants.TreasureValue
		}

		if strings.EqualFold(cellsLayout[player.CurrentPosition], constants.Hotel) {
			ownerIndex, owned := hotelOwerDict[player.CurrentPosition]

			if owned && ownerIndex != player.PlayerIndex {
				player.NetWorth -= constants.HotelRent
				playersPairs[ownerIndex].NetWorth += constants.HotelRent
			}
			if !owned && player.NetWorth >= constants.HotelWorth {
				player.NetWorth -= constants.HotelWorth
				hotelOwerDict[player.CurrentPosition] = player.PlayerIndex
			}
		}
	}

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
