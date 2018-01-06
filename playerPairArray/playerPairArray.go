package playerPairArray

import (
	"sort"
)

type PlayerPair struct {
	PlayerName string
	NetWorth   int
}
type PlayerPairArray []PlayerPair

func Sort(arrayOfPlayers PlayerPairArray) PlayerPairArray {
	sort.Sort(arrayOfPlayers)
	return arrayOfPlayers
}

func (s PlayerPairArray) Len() int {
	return len(s)
}

func (s PlayerPairArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s PlayerPairArray) Less(i, j int) bool {
	return (s[j].NetWorth) < (s[i].NetWorth)
}
