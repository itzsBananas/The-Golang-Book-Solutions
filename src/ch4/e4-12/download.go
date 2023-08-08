package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a request to
// https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
// many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd
// that, using this index, prints the URL and transcript of each comic that matches a search term
// provided on the command line.
type XKCDInfo struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

// func main() {
// 	for i := 571; i < 600; i++ {
// 		cantSleep, _ := downloadInfo(i)
// 		saveInfo(cantSleep)
// 	}
// }

func downloadInfo(id int) (*XKCDInfo, error) {
	downloadLink := fmt.Sprintf("%s%d%s", "https://xkcd.com/", id, "/info.0.json")
	fmt.Println("Downloading ", downloadLink)
	resp, err := http.Get(downloadLink)
	if err != nil {
		fmt.Println("HTTP GET Request failed")
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result XKCDInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func saveInfo(info *XKCDInfo) {
	data, err := json.MarshalIndent(*info, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	ioutil.WriteFile(fmt.Sprintf("offline/%d.json", info.Num), data, 0666)
}
