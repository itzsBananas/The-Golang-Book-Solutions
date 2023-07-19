// Exercise 7.11: Add additional handlers so that clients can create, read, update, and delete
// database entries. For example, a request of the form /update?item=socks&price=6 will
// update the price of an item in the inventory and report an error if the item does not exist or if
// the price is invalid. (Warning: this change introduces concurrent variable updates.)
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		price := req.URL.Query().Get("price")
		parsed, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "Could not parse as a price: %s\n", price)
		} else {
			fmt.Fprintf(w, "Created item %s of price $%.2f\n", item, parsed)
		}
		db[item] = dollars(float32(parsed))
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "already such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		price := req.URL.Query().Get("price")
		parsed, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "Could not parse as a price: %s\n", price)
		} else {
			fmt.Fprintf(w, "Updated item %s of price $%.2f\n", item, parsed)
		}
		db[item] = dollars(float32(parsed))
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "Deleted item %s\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
