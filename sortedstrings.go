package handy

// SortedStrings stores strings, assumed sorted.
type SortedStrings []string

// Minus returns all words in ss that is not in yass. In
// other words, it is ss - yass.
func (ss SortedStrings) Minus(yass SortedStrings) []string {
	ret := []string{}
	l, r := 0, len(yass)
	for _, word := range ss {
		// search through ss
		region := yass[l:r]
		index := region.BinarySearch(word)
		if index >= len(region) || word != region[index] {
			ret = append(ret, word)
		}
		l = index
	}

	return ret
}

// BinarySearch finds word in the sorted string list, if
// found, it returns the index, otherwise, it returns the
// index where the word can be inserted without breaking ordering.
func (ss SortedStrings) BinarySearch(word string) int {
	l, r := 0, len(ss)-1
	for l <= r {
		mid := l + (r-l)>>1
		if ss[mid] == word {
			return mid
		} else {
			if ss[mid] > word {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	// This takes a little small proof why doing so.
	return l
}
