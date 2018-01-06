package readInputs

import (
	"Business-House-Game/constants"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ReadNumberOfPlayers() (int, error) {
	var numberOfPlayers int
	var err error

	fmt.Println("Enter Number of players(min number of players allowed 2:")
	_, err = fmt.Scanf("%d", &numberOfPlayers)

	if err != nil {
		err = errors.New("Please Enter Numbers only for number of players")
	}
	if numberOfPlayers < 2 {
		err = errors.New("Min number of players allowed 2")
	}
	return numberOfPlayers, err
}

func ReadCellsPositonAndTypes() ([]string, error) {
	var cellsPositonsString string
	fmt.Println("Enter CellsType with comma separted.")
	fmt.Println("Allowerd values are E (Empty) J (Jail) T (Treasure) H (Hotel)")
	fmt.Println("Example of usage E,E,T,H")
	fmt.Scan(&cellsPositonsString)
	return SplitAndconvertToCellTypeArray(cellsPositonsString)
}

func SplitAndconvertToCellTypeArray(inputstr string) ([]string, error) {
	var cellsPositons []string
	var exception error

	inputstr = strings.Trim(inputstr, " ")
	for _, cellValue := range strings.Split(inputstr, ",") {

		if !strings.EqualFold(cellValue, constants.Empty) &&
			!strings.EqualFold(cellValue, constants.Jail) &&
			!strings.EqualFold(cellValue, constants.Treasure) &&
			!strings.EqualFold(cellValue, constants.Hotel) {
			exception = errors.New("Only Allowed values E J T H")
			break
		}

		cellsPositons = append(cellsPositons, strings.ToUpper(cellValue))
	}
	return cellsPositons, exception
}

func ReadDiceValues() ([]int, error) {
	var diceOutputStr string

	fmt.Println("Enter Dice Values separated by comma. eg 1,2,3 :")
	fmt.Scan(&diceOutputStr)
	return SplitAndConvertToIntArray(diceOutputStr)
}

func SplitAndConvertToIntArray(inputstr string) ([]int, error) {
	var exception error
	var cellsPositons []int

	inputstr = strings.Trim(inputstr, " ")
	for _, number := range strings.Split(inputstr, ",") {
		value, err := strconv.Atoi(number)
		if err != nil {
			exception = errors.New("Only numbers are allowed as dice value")
			break
		}
		cellsPositons = append(cellsPositons, value)
	}
	return cellsPositons, exception
}
