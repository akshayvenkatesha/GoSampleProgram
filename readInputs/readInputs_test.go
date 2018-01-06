package readInputs

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplitAndConverToIntArrayPositive(t *testing.T) {

	intSampleArray := []int{1, 2, 3, 4, 5, 6, 7}

	intArray, err := SplitAndConvertToIntArray(
		strings.Trim(strings.Replace(fmt.Sprint(intSampleArray), " ", ",", -1), "[]"))

	if err != nil {
		t.Errorf("Parsing failed with exception ")
		t.Error(err)
	}
	for i, value := range intArray {
		if intSampleArray[i] != value {
			t.Errorf("Expected %d but found %d", intSampleArray[i], value)
		}
	}
}

func TestSplitAndConverToIntArrayNegative(t *testing.T) {

	sampleStr := "1,2,3,a,5"
	_, err := SplitAndConvertToIntArray(sampleStr)

	if err == nil {
		t.Errorf("Expected to fail but did not faild. passed wrong input %s", sampleStr)
	}
}

func TestSplitAndconvertToCellTypeArrayPositive(t *testing.T) {

	sampleArray := []string{"T", "J", "E", "E", "t", "j", "h", "H", "e", "T", "T"}

	intArray, err := SplitAndconvertToCellTypeArray(strings.Join(sampleArray, ","))

	if err != nil {
		t.Errorf("Parsing failed with exception ")
		t.Error(err)
	}
	for i, value := range intArray {
		if !strings.EqualFold(sampleArray[i], value) {
			t.Errorf("Expected %s but found %s", sampleArray[i], value)
		}
	}
}

func TestSplitAndconvertToCellTypeArrayNegative(t *testing.T) {

	sampleArray := []string{"T", "J", "E", "E", "t", "jj", "0", "H", "e", "T", "T"}
	_, err := SplitAndconvertToCellTypeArray(strings.Join(sampleArray, ","))

	if err == nil {
		t.Errorf("Expected to fail but did not faild. passed wrong input %s", sampleArray)
	}
}
