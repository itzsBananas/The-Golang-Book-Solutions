// Exercise 7.8: Many GUIs provide a table widget with a stateful multi-tier sort: the primary
// sort key is the most recently clicked column head, the secondary sort key is the second-most
// recently clicked column head, and so on. Define an implementation of sort.Interface for
// use by such a table. Compare that approach with repeated sorting using sort.Stable.

package multi_tier_sort

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

// !+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

//!-main

// !+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-printTracks

// !+multi_tier
type multi_tier struct {
	t          []*Track
	sort_order [5]int
}

func (x multi_tier) Len() int      { return len(x.t) }
func (x multi_tier) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x multi_tier) Less(i, j int) bool {
	a, b := x.t[i], x.t[j]
	for _, column_index := range x.sort_order {
		switch column_index {
		case 1, -1:
			if a.Title != b.Title {
				return a.Title < b.Title == (column_index > 0)
			}
		case 2, -2:
			if a.Artist != b.Artist {
				return a.Artist < b.Artist == (column_index > 0)
			}
		case 3, -3:
			if a.Album != b.Album {
				return a.Length < b.Length == (column_index > 0)
			}
		case 4, -4:
			if a.Year != b.Year {
				return a.Year < b.Year == (column_index > 0)
			}
		case 5, -5:
			if a.Length != b.Length {
				return a.Length < b.Length == (column_index > 0)
			}
		}
	}
	return false
}

//!-multi_tier

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
