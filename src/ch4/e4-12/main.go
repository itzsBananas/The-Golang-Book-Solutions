package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	base   = 571
	length = 29
)

// Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a request to
// https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
// many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd
// that, using this index, prints the URL and transcript of each comic that matches a search term
// provided on the command line.
func main() {
	var info [length]*XKCDInfo
	index := make(map[string]map[int]bool)
	for i := 571; i < 600; i++ {
		data, err := indexInfo(i)
		if err != nil {
			return
		}
		info[i-base] = data

		for _, word := range strings.Split(data.Transcript, " ") {
			word = strings.TrimSuffix(word, ",")
			word = strings.TrimSuffix(word, ".")
			if index[word] == nil {
				index[word] = make(map[int]bool, 0)
			}
			index[word][i] = true
		}
	}
	// fmt.Printf("%#v\n", index)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		searchTerm := input.Text()

		matches := index[searchTerm]
		if matches == nil {
			fmt.Printf("Search term failed")
			return
		}

		fmt.Printf("%d matches found!\n", len(matches))
		for comicId := range matches {
			data := info[comicId-base]
			fmt.Printf("%s:\n%s\n", "https://xkcd.com/"+fmt.Sprintf("%d", data.Num), data.Transcript)
		}
		fmt.Printf("\n")
	}
}

func indexInfo(id int) (*XKCDInfo, error) {
	var result XKCDInfo
	file, err := os.Open(fmt.Sprintf("offline/%d.json", id))
	if err != nil {
		log.Fatal("Error opening file")
		return nil, err
	}

	if err := json.NewDecoder(file).Decode(&result); err != nil {
		log.Fatal("Error decoding file")
		return nil, err
	}
	return &result, nil
}
