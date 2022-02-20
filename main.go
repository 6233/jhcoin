package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/6233/jhcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle	string
	blocks	[]*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
	data := homeData{
		PageTitle: "home",
		blocks : blockchain.GetBlockchain().AllBlock(),
	}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}