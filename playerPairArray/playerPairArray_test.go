package playerPairArray

import "testing"

func TestSort(t *testing.T) {
	unSortedArray := PlayerPairArray{
		{"player3", -25},
		{"player1", 42},
		{"player2", 10},
		{"player4", 25},
	}
	sortedSampleArray := PlayerPairArray{
		{"player1", 42},
		{"player4", 25},
		{"player2", 10},
		{"player3", -25},
	}

	sortedArray := Sort(unSortedArray)

	for i, player := range sortedArray {
		if player.playerName != sortedSampleArray[i].playerName {
			t.Errorf("Expected %s playername but found %s playername", sortedSampleArray[i].playerName, player.playerName)
		}
		if player.netWorth != sortedSampleArray[i].netWorth {
			t.Errorf("Expected %d Networth but found %d playername", sortedSampleArray[i].netWorth, player.netWorth)
		}
	}
}
