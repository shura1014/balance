package balance

import (
	"sort"
	"testing"
)

func TestWeight_MaxCommonDivisor(t *testing.T) {

	var weightArray = []int{8, 4, 10}

	sort.Ints(weightArray)
	var result int
	if len(weightArray) > 0 {
		for i := weightArray[0]; i >= 2; i-- {
			isDivideExactly := false
			for _, w := range weightArray {
				isDivideExactly = w%i == 0
				if !isDivideExactly {
					break
				}
			}
			if isDivideExactly {
				result = i
				break
			}
		}
	}
	t.Log(result)
}
