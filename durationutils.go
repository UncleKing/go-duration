package duration

import "sort"

type Duration struct {
	from uint64
	to   uint64
}

type ByFrom []Duration

// merges the two sorted durations lists.
func MergeDurations(one, two []Duration) []Duration {
	len1 := len(one)
	len2 := len(two)

	// it will not be more than len1 + len2
	var merged []Duration

	i1, i2 := 0, 0
	m1 := true
	m2 := true
	var d2, d1 Duration
	for {
		if m1 {
			d1 = one[i1]
		}
		if m2 {
			d2 = two[i2]
		}
		if d1.from < d2.from {
			// one starts sooner
			if d1.to < d2.from {
				// one ends before two starts --> no over lap progress one
				m1 = true
				m2 = false
				merged = append(merged, d1)
			} else {
				// one ends after two starts
				if d1.to < d2.to {
					// one ends before two ends --> create new with one.from and two.end and replace that with two
					d2.from = d1.from
					m1 = true
					m2 = false
				} else {
					// one ends after two ends --> progress two keep one as is.
					m2 = true
					m1 = false
				}
			}
		} else {
			// two starts sooner
			if d2.to < d1.from {
				// two ends before one starts --> no over lap progress two
				m2 = true
				m1 = false
				merged = append(merged, d2)
			} else {
				// two ends after one starts
				if d2.to < d1.to {
					// two ends before one ends --> create new with two.from and one.end and replace that with one
					d1.from = d2.from
					m1 = true
					m2 = false
				} else {
					// two ends after one ends --> progress one keep two as is.
					m1 = true
					m2 = false
				}
			}
		}
		if m1 {
			i1++
		}

		if m2 {
			i2++
		}

		if i1 >= len1 {
			// add the remaining l2
			for i := i2; i < len2; i++ {
				merged = append(merged, two[i])
			}
			break
		}
		if i2 >= len2 {
			// add the remaining l1
			for i := i1; i < len1; i++ {
				merged = append(merged, one[i])
			}
			break
		}
	}
	return merged
}

func HasConflict(one, two []Duration) bool {
	len1 := len(one)
	len2 := len(two)

	i1, i2 := 0, 0
	m1 := true
	m2 := true
	var d2, d1 Duration
	for {
		if m1 {
			d1 = one[i1]
		}
		if m2 {
			d2 = two[i2]
		}
		if d1.from < d2.from {
			// one starts sooner
			if d1.to < d2.from {
				// one ends before two starts --> no over lap progress one
				m1 = true
				m2 = false
			} else {
				// one ends after two starts

				// one ends before two ends

				// one ends after two ends
				return true

			}
		} else {
			// two starts sooner
			if d2.to < d1.from {
				// two ends before one starts --> no over lap progress two
				m2 = true
				m1 = false
			} else {
				// two ends after one starts
				// two ends before one ends --> create new with two.from and one.end and replace that with one
				// two ends after one ends --> progress one keep two as is.
				return true

			}
		}
		if m1 {
			i1++
		}

		if m2 {
			i2++
		}

		if i1 >= len1 || i2 >= len2 {
			break
		}
	}
	return false

}

// Sorts the Duration object. This is assumed to be no
func Sort(one []Duration) {
	sort.Slice(one, func(i, j int) bool {
		return one[i].from < one[j].from
	})
}


// HasOverlap goes thru the list of durations and returns true if there's any overlap in any of the
// duration list
// It assumes that the list is sorted.
// if the list is not sorted please call sort first.
func HasOverLap(dl []Duration) bool{
	dllen := len(dl)
	for i := 1; i < dllen; i++ {
		if dl[i].from < dl[i].to {
			return false
		}
		if dl[i - 1].to >= dl[i].from {
			return false
		}
	}
	return true
}
