package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type dollers float32

func (d dollers) String() string { return fmt.Sprintf("$%.2f", d) }

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/get", db.get)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/create", db.create)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

/*
 * TODO:
 * ・適切な HTTP Method を使う
 * ・重複の削除
 */

type database map[string]dollers

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var t = template.Must(template.New("table").Parse(`
	<table>
	<tr>
	  <th>Item</th>
	  <th>Price</th>
	</tr>
	{{range $item, $price := .}}
	<tr>
	  <td>{{$item}}</td>
	  <td>{{$price}}</td>
	</tr>
	{{end}}
	<tr>
	</table>
	`))
	if err := t.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) get(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s %s\n", item, db[item])
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, "item is nil\n", item)
		return
	}
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "%q already exists\n", item)
		return
	}
	priceParam := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceParam, 32)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "%q is not float\n", priceParam)
		return
	}
	db[item] = dollers(price)
	fmt.Fprintf(w, "%s %s\n", item, db[item])
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprint(w, "success")
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	priceParam := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceParam, 32)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "%q is not float\n", priceParam)
	}
	db[item] = dollers(price)
	fmt.Fprintf(w, "%s %s\n", item, db[item])
}
