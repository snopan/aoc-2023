package day7

import (
	"sort"
	"testing"

	"gotest.tools/assert"
)

func Test_SortHands(t *testing.T) {
	hands := []string{"KK", "KJ", "K5"}
	sort.Sort(SortJokerHands(hands))
	assert.Equal(t, hands[0], "KJ")
	assert.Equal(t, hands[1], "K5")
	assert.Equal(t, hands[2], "KK")

	hands = []string{"J5", "JJ", "J5"}
	sort.Sort(SortJokerHands(hands))
	assert.Equal(t, hands[0], "JJ")
	assert.Equal(t, hands[1], "J5")
	assert.Equal(t, hands[2], "J5")
}
