// Exercise 7.12: Change the handler for /list to print its output as an HTML table, not text.
// You may find the html/template package (§4.6) useful.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

var templ = template.Must(template.New("itemlist").Parse(`
<h1>Items (Length: {{len .}})</h1>
<table>
<tr style='text-align: left'>
<th>Item</th>
<th>Price</th>
</tr>
{{range $key, $value := .}}
<tr>
<td>{{$key}}</td>
<td>{{$value}}</td>
</tr>
{{end}}
</table>
`))

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := templ.Execute(w, db); err != nil {
		log.Fatal(err)
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
