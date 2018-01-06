package main

import (
	"Business-House-Game/constants"
	"Business-House-Game/playerPairArray"
	"Business-House-Game/readInputs"
	"fmt"
	"strconv"
)

func main() {

	var NumberOfPlayers int
	var CellsLayout []string
	var DiceOutArray []int

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

	DiceOutArray, errDiceRead := readInputs.ReadDiceValues()
	if errDiceRead != nil {
		fmt.Println(errDiceRead)
		return
	}
	fmt.Println(DiceOutArray)

	InitializePlayers(NumberOfPlayers)
}

func InitializePlayers(NumberOfPlayers int) playerPairArray.PlayerPairArray {
	var playerArray playerPairArray.PlayerPairArray

	for i := 1; i <= NumberOfPlayers; i++ {
		playerPair := playerPairArray.PlayerPair{PlayerName: "Player-" + strconv.Itoa(i), NetWorth: constants.InitialMoney}
		playerArray = append(playerArray, playerPair)
	}
	return playerArray
}
