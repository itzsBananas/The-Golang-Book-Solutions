package multi_tier_sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func BenchmarkMultiTier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tracks = []*Track{
			{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		}
		var sort_order [5]int
		for i := 0; i < len(sort_order); i++ {
			sort_order[i] = getRandomValidNum()
		}

		unsorted := multi_tier{tracks, sort_order}
		sort.Sort(unsorted)
	}
}

func getRandomValidNum() int {
	return rand.Intn(11) - 5
}
