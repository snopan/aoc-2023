package day7

import "unicode"

type SortHands []string

var letterRanking = map[rune]int{
	'T': 1,
	'J': 2,
	'Q': 3,
	'K': 4,
	'A': 5,
}

func (s SortHands) Len() int           { return len(s) }
func (s SortHands) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortHands) Less(i, j int) bool { return HandsLess(s, i, j, false) }

type SortJokerHands []string

func (s SortJokerHands) Len() int           { return len(s) }
func (s SortJokerHands) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortJokerHands) Less(i, j int) bool { return HandsLess(s, i, j, true) }

func HandsLess(s []string, i, j int, handleJoker bool) bool {
	handI := s[i]
	handJ := s[j]

	for index, ci := range handI {
		ji := rune(handJ[index])
		isCIJoker := ci == 'J'
		isJIJoker := ji == 'J'
		isCINumber := unicode.IsDigit(ci)
		isJINumber := unicode.IsDigit(ji)

		if handleJoker && (isCIJoker || isJIJoker) {
			if isCIJoker && isJIJoker {
				continue
			}
			return isCIJoker
		}
		if isCINumber && isJINumber {
			ciNumber := int(ci - '0')
			jiNumber := int(ji - '0')
			if ciNumber == jiNumber {
				continue
			}
			return ciNumber < jiNumber
		} else if !isCINumber && !isJINumber {
			ciLetterRank := letterRanking[ci]
			jiLetterRank := letterRanking[ji]
			if ciLetterRank == jiLetterRank {
				continue
			}
			return ciLetterRank < jiLetterRank
		} else {
			return isCINumber && !isJINumber
		}
	}
	return false
}
